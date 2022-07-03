package mongo

import (
	"context"
	"fmt"
	"relay_server/pkg/logger"
	"testing"
	"time"
	"os"

	cnf "relay_server/pkg/config"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

)

type fakeLog struct{}

func (t *fakeLog) Info(s interface{})  {}
func (t *fakeLog) Error(s interface{}) {}
func (t *fakeLog) Debug(s interface{}) {
	fmt.Println(s)
}

var fakeLog_test logger.ILogger= (*fakeLog)(nil)

func Test_NewClient_Connection(t *testing.T) {
	os.Setenv(cnf.ConfMongoPort, "27017")
	os.Setenv(cnf.ConfMongoHost, "localhost")
	os.Setenv(cnf.ConfMongoDBName, "radmir")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ctxDisconn, disconn := context.WithCancel(context.Background())
	
	database, err := New(ctx, ctxDisconn, &fakeLog{})
	assert.True(t, err == nil)
	assert.True(t, database != nil)
	collections, err := database.ListCollectionNames(ctx, bson.M{})
	assert.True(t, err == nil)
	fmt.Println(collections)
	
	disconn()
	time.Sleep(1*time.Second)
}