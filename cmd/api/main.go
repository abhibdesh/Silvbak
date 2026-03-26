package main

import (
	"context"
	"fmt"
	"log"
	"silvbak/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// load config
	mongoConfig := config.Load()
	client, err := mongo.Connect(
		options.Client().ApplyURI(mongoConfig.MongoURI),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = client.Ping(ctx, nil); err != nil {
		log.Fatal("Mongo not connected:", err)
	}

	fmt.Println("Mongo connected")
	db := client.Database(mongoConfig.DBName)
	fmt.Println("Using DB:", db.Name())
	// init DB
	// init services
	// start server

}
