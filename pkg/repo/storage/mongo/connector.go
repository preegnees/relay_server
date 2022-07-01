package mongo

import (
	"context"
	"fmt"
	"relay_server/pkg/logger"
	"sync"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type CnfMongo struct {
	Log      logger.ILogger
	Ctx      context.Context
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

var (
	once sync.Once
	err  error
	db   *mongo.Database
)

const (
	ErrorNewClientMongo = "$Ошибка создания клиента базы данных mongodb"
	ErrorConnectMongo   = "$Ошибка подключения к mongodb"
	ErrorPingMongo = "$Ошибка при пинге mongodb"
	
	ErrorGetDatabases = "$Ошибка получения списка баз данных"
)

func NewClient(cnf CnfMongo) (*mongo.Database, error) {
	once.Do(func() {
		var URI string
		if cnf.Username == "" && cnf.Password == "" {
			URI = fmt.Sprintf("mongodb://%s:%s", cnf.Host, cnf.Port)
		} else {
			URI = fmt.Sprintf("mongodb://%s:%s@%s:%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port)
		}
		cnf.Log.Debug(fmt.Sprintf("URL mongo: %s", URI))

		credential := options.Credential{
			Username: cnf.Username,
			Password: cnf.Password,
		}
		client, err := mongo.NewClient(options.Client().ApplyURI(URI).SetAuth(credential))
		if err != nil {
			err = fmt.Errorf("%s. Err: %v", ErrorNewClientMongo, err)
			db = nil
			return
		}

		cnf.Log.Debug(fmt.Sprintf("Успешное создание клиента"))

		err = client.Connect(cnf.Ctx)
		if err != nil {
			err = fmt.Errorf("%s. Err: %v", ErrorConnectMongo, err)
			db = nil
			return
		}

		cnf.Log.Debug(fmt.Sprintf("Успешное подключение"))

		defer client.Disconnect(cnf.Ctx)
		defer cnf.Log.Debug(fmt.Sprintf("Успешное отключение"))

		err = client.Ping(cnf.Ctx, readpref.Primary())
		if err != nil {
			err = fmt.Errorf("%s. Err: %v", ErrorPingMongo, err)
			db = nil
			return
		}
		databases, err := client.ListDatabaseNames(cnf.Ctx, bson.M{})
		if err != nil {
			err = fmt.Errorf("%s. Err: %v", ErrorGetDatabases, err)
			db = nil
			return
		}
		fmt.Println(databases)
	})
	return db, err
}
