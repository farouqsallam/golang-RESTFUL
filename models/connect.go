package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client


func ConnectDB() {
	var err error

	URI := "mongodb://localhost:27017"
	Client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil{
		panic(err)
	}

	fmt.Println("DB connected")
}