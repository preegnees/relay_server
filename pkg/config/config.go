package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type config struct {
	Port int
}

func Get() (*config, error) {
	err := godotenv.Load("config.env")
	if err != nil {
		return nil, fmt.Errorf("$Невозможно прочитать файл конфигурации. err: %v", err)
	}
	port, err := strconv.Atoi(os.Getenv("serverPort"))
	if err != nil {
		return nil, fmt.Errorf("$Невалидный порт. err: %v", err)
	}
	cnf := config{
		Port: port,
	}
	return &cnf, nil
}
