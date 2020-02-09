package database

// import (
// 	"context"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"go.mongodb.org/mongo-driver/mongo/readpref"
// )

// Connect returns mongoDB client for queries
// func Connect(uri string) *mongo.Client {
// 	options := options.Client().ApplyURI(uri)
// 	client, err := mongo.NewClient(options)

// 	if err != nil {
// 		panic(err)
// 	}

// 	client.Connect(context.Background())

// 	err = client.Ping(context.Background(), readpref.Primary())

// 	if err != nil {
// 		panic(err)
// 	}

// 	return client
// }
