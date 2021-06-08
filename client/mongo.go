package client

import (
	"context"
	"sync"
	"time"
	"url-shortener/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var mongoError error
var mongoInit sync.Once

// GetMongoClient Returns a reference to a mongo.Client
func GetMongoClient() (*mongo.Client, error) {
	mongoInit.Do(func() {
		mongoConfig := config.GetMongoConfig()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		clientOptions := options.Client().ApplyURI(mongoConfig.ConnectionString)

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			mongoError = err
		}

		err = client.Ping(ctx, nil)
		if err != nil {
			mongoError = err
		}

		mongoClient = client
	})

	return mongoClient, mongoError
}
