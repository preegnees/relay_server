package config

import (
	"os"
	"strings"
	"testing"
	"fmt"

	"github.com/stretchr/testify/assert"

	"relay_server/pkg/logger"
)

//%%% Тут нужно узнать, правда ли конфиг будте считываться только один раз

// Хранит состояние, колличество вызовов
func state() func(bool) int {
	count := 0
	return func(c bool) int {
		if c {
			count++
		}
		return count
	}
}

var st func(bool) int = state()

// фейковая структура логгера, где мы перехватываем сообщение при дебаге
type fakeLog struct{}

func (t *fakeLog) Info(s interface{})  {}
func (t *fakeLog) Error(s interface{}) {}
func (t *fakeLog) Debug(s interface{}) {
	if s == DbgConfigHasBeenRead {
		st(true)
	}
	if s == DbgAConfigRequestHasBeenMade {
		st(false)
	}
	fmt.Println(s)
}

// Тут проверяется соответсвие фейковой структуры на
var lg logger.ILogger = (*fakeLog)(nil)

// фейковая загрузка данных в перменные окружения из файла
func fakeLoadEnv1() error {
	os.Setenv(ConfServerPort, "9999")
	os.Setenv(ConfMongoPort, "8888")
	os.Setenv(ConfMongoHost, "localhost")
	os.Setenv(ConfMongoDBName, "radmir")
	return nil
}

// Тестируется вариант, когда у конфиг читается несколько раз.
// Должно: один раз прочититать конфиг, потом отдавать структуру
func Test_Get_OneRead(t *testing.T) {
	log := &fakeLog{}
	_ = Get(log, fakeLoadEnv1)
	_ = Get(log, fakeLoadEnv1)
	_ = Get(log, fakeLoadEnv1)
	count := st(false)
	assert.Equal(t, 1, count)
	assert.True(t, os.Getenv("server.port") == "9999")
}

//%%% Должны возникать ошибки, если параметры кофига будут невалидны

// фейковая загрузка данных в перменные окружения из файла
func fakeLoadEnv2() error {
	os.Setenv(ConfServerPort, "9999ф")
	return nil
}

// Должна быть ошибка если мы закрузим в переменные среды не число
func Test_Get_ErrorOnInvalidConfig(t *testing.T) {
	log := &fakeLog{}
	err := Get(log, fakeLoadEnv2)
	assert.True(t, err != nil)
	assert.True(t, strings.Contains(err.Error(), ErrorInvalidServerPort))
	assert.False(t, strings.Contains(err.Error(), "hello world"))
}
