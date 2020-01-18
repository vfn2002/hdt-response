package database

import (
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect returns mongoDB client for queries
func Connect(uri string) *mongo.Client {
	options := options.Client().ApplyURI(uri)
	//TODO: Add error handling
	client, _ := mongo.NewClient(options)

	client.Connect(context.Background())
	//TODO: Add error handling "err :="
	client.Ping(context.Background(), readpref.Primary())

	return client
}
