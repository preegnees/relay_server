package config

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"

	"relay_server/pkg/logger"
)

const (
	EnvFileName                        = "config.env"
	ErrorUnableToReadConfigurationFile = "$Невозможно прочитать файл конфигурации"

	ConfServerPort         = "server.port"
	ErrorInvalidServerPort = "$Невалидный порт сервера"
	ConfMongoPort         = "mongo.port"
	ErrorInvalidMonogPort = "$Невалидный порт моногодб"
	ConfMongoHost         = "mongo.host"
	ErrorInvalidMongoHost = "$Невалидный хост моногодб"
	ConfMongoUsername         = "mongo.username"
	ErrorInvalidMongoUsername = "$Невалидный username моногодб"
	ConfMongoPassword         = "mongo.password"
	ErrorInvalidMongoPassword = "$Невалидный password моногодб"
	ConfMongoDBName         = "mongo.db"
	ErrorInvalidMongoDBName = "$Невалидное имя базы данных"
)

const (
	DbgConfigHasBeenRead         = "Конфиг был прочтен"
	DbgAConfigRequestHasBeenMade = "Был сделан запрос конфига"
)

type serverConf struct {
	Port int
}

type mongoConf struct {
	Port     int
	Host     string
	Username string
	Password string
	DB       string
}

type config struct {
	serverConf
	mongoConf
}

var (
	cnf  config
	e    error
	once sync.Once
)

func LoadEnv() error {
	err := godotenv.Load(EnvFileName)
	if err != nil {
		return fmt.Errorf("%s. err: %v", ErrorUnableToReadConfigurationFile, err)
	}
	return nil
}

func Get(log logger.ILogger, loadEnv func() error) (*config, error) {
	once.Do(func() {
		err := loadEnv()
		if err != nil {
			e = err
			cnf = config{}
			return
		}
		// server Port
		serverPort, err := strconv.Atoi(os.Getenv(ConfServerPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidServerPort, err)
			cnf = config{}
			return
		}
		// mongo Port
		mongoPort, err := strconv.Atoi(os.Getenv(ConfMongoPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMonogPort, err)
			cnf = config{}
			return
		}
		// mongo Host
		mongoHost := os.Getenv(ConfMongoHost)
		if mongoHost == "" {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMongoHost, err)
			cnf = config{}
			return
		}
		// mongo Username
		mongoUsername := os.Getenv(ConfMongoUsername)
		// mongo Password
		mongoPassword := os.Getenv(ConfMongoPassword)
		// mongo db
		db := os.Getenv(ConfMongoDBName)
		if db == "" {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMongoDBName, err)
			cnf = config{}
			return
		}
		cnf = config{
			serverConf: serverConf{
				Port: serverPort,
			},
			mongoConf: mongoConf{
				Port: mongoPort,
				Host: mongoHost,
				Username: mongoUsername,
				Password: mongoPassword,
				DB: db,
			},
		}
		log.Debug(cnf)

		log.Debug(DbgConfigHasBeenRead)
	})
	log.Debug(DbgAConfigRequestHasBeenMade)
	return &cnf, e
}
