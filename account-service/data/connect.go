package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(ctx context.Context, url string) (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Println("Error connecting to mongo")
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("Error pinging mongo")
		log.Fatal(err)
	}

	return client, err
}
