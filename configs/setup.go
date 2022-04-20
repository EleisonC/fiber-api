package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://chrisgolang:Jo0BS8UJdYSSWTCh@cluster0-shard-00-00.vjnxv.mongodb.net:27017,cluster0-shard-00-01.vjnxv.mongodb.net:27017,cluster0-shard-00-02.vjnxv.mongodb.net:27017/golangDB?ssl=true&replicaSet=atlas-5jq5ul-shard-0&authSource=admin&retryWrites=true&w=majority"))
	fmt.Println(err)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println(ctx)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client
}

var DB *mongo.Client = ConnectDB()

// getting database collections

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("golangAPI").Collection((collectionName))
	return collection
}
