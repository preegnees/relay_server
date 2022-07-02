package mongo

import (
	"context"
	"sync"
	"fmt"

	"relay_server/pkg/logger"
	repo "relay_server/pkg/repo/storage"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	ErrorNewClient = "$ОШшибка при создании клиента mongo"
)

type mongoDB struct {
	Database *mongo.Database
}

var (
	db *mongoDB
	once sync.Once
	e error
)

var mongoTest repo.IInfoHost = (*mongoDB)(nil)

func (m *mongoDB) GetAllInfoHost(ctx context.Context, log logger.ILogger) (*[]repo.Host, error) { 
	_ = m.Database.Collection("hello")
	return &[]repo.Host{}, nil
}

func (m *mongoDB) GetIdInfoHost(ctx context.Context, log logger.ILogger, id string) (*repo.Host, error) { 
	
	return &repo.Host{}, nil
}

func (m *mongoDB) SaveInfoHost(ctx context.Context, h *repo.Host) error {
	return nil
}

func NewStorage() *mongoDB {
	once.Do(func(){
		cli, err := NewClient(cnf)
		if err != nil {
			e = fmt.Errorf("%s. Err: %v", ErrorNewClient, err)
		}
		db = &mongoDB{
			Database: cli.Database(cnf.DB),
		}
	})
	return db
}