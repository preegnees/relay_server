package mongo

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type fakeLog struct{}

func (t *fakeLog) Info(s string)  {}
func (t *fakeLog) Error(s string) {}
func (t *fakeLog) Debug(s string) {}

func Test_NewClient_Connection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var cnf *CnfMongo = &CnfMongo{
		Log:      &fakeLog{},
		Ctx:      ctx,
		Host:     "localhost",
		Port:     "8888",
		Username: "",
		Password: "",
		DB:       "test_db",
	}

	database, err := NewClient(cnf)
	assert.True(t, err == nil)
	assert.True(t, database != nil)
	cancel()
}