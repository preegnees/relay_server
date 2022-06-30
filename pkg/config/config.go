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

type config struct {
	Port int
}

var (
	cnf  config
	e    error
	once sync.Once
)

func Get(log logger.ILogger) (*config, error) {
	once.Do(func() {
		err := godotenv.Load(EnvFileName)
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorUnableToReadConfigurationFile, err)
		}
		port, err := strconv.Atoi(os.Getenv(ConfServerPort))
		if err != nil {
			e = fmt.Errorf("%s. err: %v", ErrorInvalidServerPort, err)
		}
		cnf = config{
			Port: port,
		}
		log.Debug("Конфиг был прочтен")
	})
	log.Debug("Был сделан запрос конфига")
	return &cnf, e
}
