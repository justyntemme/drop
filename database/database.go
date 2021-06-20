package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DB() *mongo.Client {
	credential := options.Credential{
		Username: "admin",
		Password: "nKH.=XdYp#-ECw,=gW",
	}
	clientOpts := options.Client().ApplyURI("mongodb://144.202.66.168:27018").SetAuth(credential)

	// Connect to MongoDB

	client, err := mongo.Connect(context.TODO(), clientOpts)
	if err != nil {
		log.Fatal(err)
	}
	_ = client

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
