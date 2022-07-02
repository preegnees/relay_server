package mongo

import (
	"context"
	"fmt"
	"testing"
	"time"

	repo "relay_server/pkg/repo/storage"

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
	var cnf *CnfMongo = &CnfMongo{
		Log:      &fakeLog{},
		Ctx:      ctx,
		Host:     "localhost",
		Port:     "27017",
		Username: "",
		Password: "",
		DB:       "radmir",
	}

	var storage repo.IInfoHost = NewStorage(cnf)
	storage.
}