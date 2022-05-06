package databaseUtils

import (
	"fmt"

	"time"

	"context"

	// mongodb packages

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// global variables
type MongoDbConObjectType struct {
	Client *mongo.Client
}

type MongoClient *mongo.Client

var MongoDbConObj MongoDbConObjectType

func ConnectToMongoDB(uri string) (*mongo.Client, context.Context, context.CancelFunc, error) {

	// ctx will be used to set deadline for process, here
	// deadline will of 30 seconds.
	ctx, cancel := context.WithTimeout(context.Background(),
		30*time.Second)

	// Create a new client and connect to the server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	MongoDbConObj.Client = client
	fmt.Println("Successfully connected and pinged.")

	return client, ctx, cancel, err
}

func CloseDbConnection(cancel context.CancelFunc, DbClient *mongo.Client, ctx context.Context) {
	// client provides a method to close
	// a mongoDB connection.
	defer cancel()
	defer func() {
		// client.Disconnect method also has deadline.
		// returns error if any,
		if err := DbClient.Disconnect(ctx); err != nil {
			panic(err)
		}
		fmt.Println("MongoDB connection terminated.")

	}()
}
