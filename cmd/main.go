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
	cnf, err := config.Get(log, config.LoadEnv)
	if err != nil {
		return fmt.Errorf("%s. Err: %v", ErrorReadConfig, err)
	}
	if cnf == nil {
		return fmt.Errorf("%s. Cnf: %v", ErrorInvalidConfig, cnf)
	}
	log.Debug(fmt.Sprintf("Cnf: %v", cnf))
	return nil
}

// структура для описания флагов
type Flags struct {
	Debug bool
}

// парсер флагов
func parseFlags() *Flags {
	var dbg = flag.Bool("d", false, "is debug???")
	flag.Parse()
	flgs := &Flags{
		Debug: *dbg,
	}
	return flgs
}

// инициализация логгер, атребут дебаг
func initLogger(flgs *Flags) logger.ILogger {
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