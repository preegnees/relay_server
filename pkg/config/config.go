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

	ConfServerPort         = "serverPort"
	ErrorInvalidServerPort = "$Невалидный порт"
)

const (
	DbgConfigHasBeenRead = "Конфиг был прочтен"
	DbgAConfigRequestHasBeenMade = "Был сделан запрос конфига"
)

type config struct {
	Port int
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
		}
		port, err := strconv.Atoi(os.Getenv(ConfServerPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidServerPort, err)
		}
		cnf = config{
			Port: port,
		}
		log.Debug(DbgConfigHasBeenRead)
	})
	log.Debug(DbgAConfigRequestHasBeenMade)
	return &cnf, e
}
