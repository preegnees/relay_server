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

	ConfServerPort            = "server.port"
	ErrorInvalidServerPort    = "$Невалидный порт сервера"
	ConfMongoPort             = "mongo.port"
	ErrorInvalidMonogPort     = "$Невалидный порт моногодб"
	ConfMongoHost             = "mongo.host"
	ErrorInvalidMongoHost     = "$Невалидный хост моногодб"
	ConfMongoUsername         = "mongo.username"
	ErrorInvalidMongoUsername = "$Невалидный username моногодб"
	ConfMongoPassword         = "mongo.password"
	ErrorInvalidMongoPassword = "$Невалидный password моногодб"
	ConfMongoDBName           = "mongo.db"
	ErrorInvalidMongoDBName   = "$Невалидное имя базы данных"
)

const (
	DbgConfigHasBeenRead         = "Конфиг был прочтен"
	DbgAConfigRequestHasBeenMade = "Был сделан запрос конфига"
)

func LoadEnv() error {
	err := godotenv.Load(EnvFileName)
	if err != nil {
		return fmt.Errorf("%s. err: %v", ErrorUnableToReadConfigurationFile, err)
	}
	return nil
}

var (
	e    error
	once sync.Once
)

func Get(log logger.ILogger, loadEnv func() error) error {
	once.Do(func() {
		err := loadEnv()
		if err != nil {
			e = err
			return
		}
		// server Port
		serverPort, err := strconv.Atoi(os.Getenv(ConfServerPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidServerPort, err)
			return
		}
		// mongo Port
		mongoPort, err := strconv.Atoi(os.Getenv(ConfMongoPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMonogPort, err)
			return
		}
		// mongo Host
		mongoHost := os.Getenv(ConfMongoHost)
		if mongoHost == "" {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMongoHost, err)
			return
		}
		// mongo Username
		mongoUsername := os.Getenv(ConfMongoUsername)
		// mongo Password
		mongoPassword := os.Getenv(ConfMongoPassword)
		// mongo db
		mongoDBName := os.Getenv(ConfMongoDBName)
		if mongoDBName == "" {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidMongoDBName, err)
			return
		}

		log.Debug(fmt.Sprintf(
			"%s:%v\n%s:%v\n%s:%v\n%s:%v\n%s:%v\n%s:%v\n",
			ConfServerPort, serverPort,
			ConfMongoPort, mongoPort,
			ConfMongoHost, mongoHost,
			ConfMongoUsername, mongoUsername,
			ConfMongoPassword, mongoPassword,
			ConfMongoDBName, mongoDBName,
		))

		log.Debug(DbgConfigHasBeenRead)
	})
	log.Debug(DbgAConfigRequestHasBeenMade)
	return e
}
