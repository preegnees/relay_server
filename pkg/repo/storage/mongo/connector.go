package mongo

import (
	"context"
	"fmt"
	"relay_server/pkg/logger"

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

const (
	ErrorNewClientMongo = "$Ошибка создания клиента базы данных mongodb"
	ErrorConnectMongo   = "$Ошибка подключения к mongodb"
	ErrorPingMongo      = "$Ошибка при пинге mongodb"

	ErrorGetDatabase = "$Ошибка получения базы данных"
)

func NewClient(cnf *CnfMongo) (*mongo.Client, error) {
	var URI string
	var credential options.Credential
	var client *mongo.Client 
	var err error

	if cnf.Username == "" && cnf.Password == "" {
		URI = fmt.Sprintf("mongodb://%s:%s", cnf.Host, cnf.Port)
		client, err = mongo.NewClient(options.Client().ApplyURI(URI))
	} else {
		URI = fmt.Sprintf("mongodb://%s:%s@%s:%s", cnf.Username, cnf.Password, cnf.Host, cnf.Port)
		credential = options.Credential{
			Username: cnf.Username,
			Password: cnf.Password,
		}
		client, err = mongo.NewClient(options.Client().ApplyURI(URI).SetAuth(credential))
	}
	
	cnf.Log.Debug(fmt.Sprintf("URL mongo: %s", URI))

	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorNewClientMongo, err)
	}

	cnf.Log.Debug(fmt.Sprintf("Успешное создание клиента"))

	err = client.Connect(cnf.Ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorConnectMongo, err)
	}

	cnf.Log.Debug(fmt.Sprintf("Успешное подключение"))

	// defer client.Disconnect(cnf.Ctx)
	// defer cnf.Log.Debug(fmt.Sprintf("Успешное отключение"))

	err = client.Ping(cnf.Ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorPingMongo, err)
	}

	cnf.Log.Debug(fmt.Sprintf("Успешный пинг"))

	// database := client.Database(cnf.DB)
	// if err != nil {
	// 	return nil, fmt.Errorf("%s. Err: %v", ErrorGetDatabase, err)
	// }
	return client, nil
}
