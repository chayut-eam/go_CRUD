package db

import (
	"context"
	"time"

	"crud/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	log "crud/logger"
)

var client *mongo.Client
var database *mongo.Database
var ctx = context.TODO()

func Init(config model.MongoConfig) error {
	clientOptions := options.Client()
	clientOptions.ApplyURI(config.URI)
	clientOptions.SetMinPoolSize(config.MinPoolSize)
	clientOptions.SetMaxPoolSize(config.MaxPoolSize)
	clientOptions.SetMaxConnIdleTime(time.Millisecond * time.Duration(config.MaxIdleTime))

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Logger().Fatal(err)
	}

	database = client.Database(config.Database)
	return Ping()
}

func Ping() error {
	return database.Client().Ping(ctx, nil)
}

func TearDown() {
	log.Logger().Info("Tearing down database connection...")
	if err := client.Disconnect(ctx); err != nil {
		panic(err)
	}
}

func PersonCollection() *mongo.Collection {
	return database.Collection("Person")
}
