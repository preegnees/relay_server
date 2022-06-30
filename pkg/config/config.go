package config

import (
	fmt "fmt"
	os "os"
	strconv "strconv"

	godotenv "github.com/joho/godotenv" 
)
 
type config struct {
	Port int8
}

func Get() (*config, error) {
	err := godotenv.Load("config.env")
	if err != nil {
		return nil, fmt.Errorf("$Невозможно прочитать файл конфигурации. err: %v", err)
	}
	port, err := strconv.ParseInt(os.Getenv("port"), 10, 8)
	if err != nil {
		return nil, fmt.Errorf("$Невалидный порт. err: %v", err)
	}
	cnf := config{
		Port: int8(port),
	}
	return &cnf, nil
}