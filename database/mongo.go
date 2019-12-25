package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect returns mongoDB client for queries
func Connect(uri string) *mongo.Client {
	options := options.Client().ApplyURI(uri)
	client, _ := mongo.NewClient(options)

	client.Connect(context.Background())

	return client
}
