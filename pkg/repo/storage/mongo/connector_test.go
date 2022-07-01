package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

type fakeLog struct{}

func (t *fakeLog) Info(s interface{})  {}
func (t *fakeLog) Error(s interface{}) {}
func (t *fakeLog) Debug(s interface{}) {
	fmt.Println(s)
}

func Test_NewClient_Connection(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	var cnf *CnfMongo = &CnfMongo{
		Log:      &fakeLog{},
		Ctx:      ctx,
		Host:     "localhost",
		Port:     "27017",
		Username: "",
		Password: "",
		DB:       "radmir",
	}

	client, err := NewClient(cnf)
	assert.True(t, err == nil)
	assert.True(t, client != nil)

	database, err := client.ListDatabaseNames(cnf.Ctx, bson.M{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(database)
	defer cancel()
	defer client.Disconnect(cnf.Ctx)
}