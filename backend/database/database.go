package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"lilien-server/config"
)

var databaseClient *mongo.Client

func Initlization() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	databaseClient, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().GetString("db.mongouri")))
	if err != nil {
		log.Fatal(err)
	}

	err = databaseClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
}

func GetDatabaseClient() *mongo.Client {
	return databaseClient
}
