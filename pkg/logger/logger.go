package logger

import (
	fmt "fmt"
	log "log"
)

type ILogger interface {
	Info(string)
	Debug(string)
	Error(string)	
}

type MyLogger struct {
	Dbg bool
}

func (mlg *MyLogger) Info(s string) {
	log.Println(fmt.Sprintf("INFO. %s", s))
}

func (mlg *MyLogger) Error(s string) {
	log.Fatal(fmt.Sprintf("ERROR. %s", s))
}

func (mlg *MyLogger) Debug(s string) {
	if mlg.Dbg {
		log.Println(fmt.Sprintf("DEBUG. %s", s))
	}
}