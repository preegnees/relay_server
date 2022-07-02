package mongo

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"relay_server/pkg/logger"
	"syscall"

	cnf "relay_server/pkg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	ErrorNewClientMongo = "$Ошибка создания клиента базы данных mongodb"
	ErrorConnectMongo   = "$Ошибка подключения к mongodb"
	ErrorPingMongo      = "$Ошибка при пинге mongodb"

	ErrorGetDatabase = "$Ошибка получения базы данных"
)

func New(ctx context.Context, ctxDisconn context.Context, log logger.ILogger) (*mongo.Database, error) {
	var URI string
	var credential options.Credential
	var client *mongo.Client
	var err error

	if os.Getenv(cnf.ConfMongoUsername) == "" || os.Getenv(cnf.ConfMongoPassword) == "" {
		URI = fmt.Sprintf("mongodb://%s:%s",
			os.Getenv(cnf.ConfMongoHost),
			os.Getenv(cnf.ConfMongoPort))
		client, err = mongo.NewClient(options.Client().ApplyURI(URI))
	} else {
		URI = fmt.Sprintf("mongodb://%s:%s@%s:%s",
			os.Getenv(cnf.ConfMongoUsername),
			os.Getenv(cnf.ConfMongoPassword),
			os.Getenv(cnf.ConfMongoHost),
			os.Getenv(cnf.ConfMongoPort))
		credential = options.Credential{
			Username: os.Getenv(cnf.ConfMongoUsername),
			Password: os.Getenv(cnf.ConfMongoPassword),
		}
		client, err = mongo.NewClient(options.Client().ApplyURI(URI).SetAuth(credential))
	}

	log.Debug(fmt.Sprintf("URL mongo: %s", URI))

	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorNewClientMongo, err)
	}

	log.Debug(fmt.Sprintf("Успешное создание клиента"))

	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorConnectMongo, err)
	}

	log.Debug(fmt.Sprintf("Успешное подключение к базе данных"))
	log.Info("Успешное подключение к базе данных")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("%s. Err: %v", ErrorPingMongo, err)
	}

	log.Debug(fmt.Sprintf("Успешный пинг"))

	database := client.Database(os.Getenv(cnf.ConfMongoDBName))

	go func(ctxDisconn context.Context, cli *mongo.Client) {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

		ctx := context.Background()

		disconn := func() {
			cli.Disconnect(ctx)
			log.Debug("Отлючено от базы данных")
		}
		
		select {
		case <-ctxDisconn.Done():
			disconn()
		case <-quit:
			disconn()
		}
	}(ctxDisconn, client)

	return database, nil
}
