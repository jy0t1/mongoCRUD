package main

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func db() *mongo.Client {
	db, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://test_admin:test123@clusterjs-u8ha7.mongodb.net/library_db?retryWrites=true&w=majority"))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = db.Connect(ctx)
	if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MongoDB!")

	return db
}