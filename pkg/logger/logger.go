package logger

import (
	"fmt"
	"log"
)

type ILogger interface {
	Info(interface{})
	Debug(interface{})
	Error(interface{})	
}

type MyLogger struct {
	Dbg bool
}

func (mlg *MyLogger) Info(s interface{}) {
	log.Println(fmt.Sprintf("INFO. %s", s))
}

func (mlg *MyLogger) Error(s interface{}) {
	log.Fatal(fmt.Sprintf("ERROR. %s", s))
}

func (mlg *MyLogger) Debug(s interface{}) {
	if mlg.Dbg {
		log.Println(fmt.Sprintf("DEBUG. %s", s))
	}
}

var test ILogger = (*MyLogger)(nil)