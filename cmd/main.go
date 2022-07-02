package main

import (
	"fmt"
	"flag"

	"relay_server/pkg/config"
	"relay_server/pkg/logger"
)

const (
	ErrorMainErr = "$Критическая ошибка"
	ErrorReadConfig = "$Ошибка при чтении конфига"
	ErrorInvalidConfig = "$Невалидный конфиг"
)

// точка входа
func run(log logger.ILogger) error {
	err := config.Get(log, config.LoadEnv)
	if err != nil {
		return fmt.Errorf("%s. Err: %v", ErrorReadConfig, err)
	}
	return nil
}

// структура для описания флагов
type flags struct {
	Debug bool
}

// парсер флагов
func parseFlags() *flags {
	var dbg = flag.Bool("d", false, "is debug???")
	flag.Parse()
	flgs := &flags{
		Debug: *dbg,
	}
	return flgs
}

// инициализация логгер, атребут дебаг
func initLogger(flgs *flags) logger.ILogger {
	var log logger.ILogger = &logger.MyLogger{
		Dbg: flgs.Debug,
	}
	return log
}

func main() {
	flgs := parseFlags()
	log := initLogger(flgs)
	log.Debug("Старт")
	if err := run(log); err != nil {
		log.Error(fmt.Sprintf("%s. Err: %v", ErrorMainErr, err))
	}
}