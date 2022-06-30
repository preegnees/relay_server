package main

import (
	"fmt"
	"flag"

	"relay_server/pkg/config"
	"relay_server/pkg/logger"
)

// точка входа
func run(log logger.ILogger) error {
	cnf, err := config.Get()
	if err != nil {
		return fmt.Errorf("$Ошибка при чтении конфига. Err: %v", err)
	}
	if cnf == nil {
		return fmt.Errorf("$Невалидный конфиг. Cnf: %v", cnf)
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
		log.Error(fmt.Sprintf("$Критическая ошибка. Err: %v", err))
	}
}