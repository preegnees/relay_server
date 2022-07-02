package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"
	"os"

	cnf "relay_server/pkg/config"
	// repo "relay_server/pkg/repo/storage"

	// "github.com/stretchr/testify/assert"
	// "go.mongodb.org/mongo-driver/bson"
)

type fakeLog1 struct{}

func (t *fakeLog1) Info(s interface{})  {}
func (t *fakeLog1) Error(s interface{}) {}
func (t *fakeLog1) Debug(s interface{}) {
	fmt.Println(s)
}

func Test_NewStorage_Create(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	ctxDisconn, _:= context.WithCancel(context.Background())

	os.Setenv(cnf.ConfMongoPort, "27017")
	os.Setenv(cnf.ConfMongoHost, "localhost")
	os.Setenv(cnf.ConfMongoDBName, "radmir")

	_, _ = New(ctx, ctxDisconn, &fakeLog1{})
}