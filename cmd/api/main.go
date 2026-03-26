package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {
	// load config
	// init DB
	// init services
	// start server
	client, err := mongo.Connect(
		options.Client().ApplyURI("DB_LINK"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ping to verify connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Mongo not connected:", err)
	}

	fmt.Println("Mongo connected ✅")

	db := client.Database("silvbak")
	fmt.Println("Using DB:", db.Name())
	fmt.Println("sdf", client)
	fmt.Println("sdf", err)
}
