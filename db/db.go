package db

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://<username>:<password>@<endpoint>/?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	return client
}
