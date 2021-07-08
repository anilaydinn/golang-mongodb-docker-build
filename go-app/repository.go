package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	client *mongo.Client
}

type DataEntity struct {
	ID   string `bson:"id"`
	Text string `bson:"text"`
}

func NewRepository() *Repository {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://mongodb:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return &Repository{client}
}
