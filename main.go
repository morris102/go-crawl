package main

import (
	"context"
	"fmt"
	"go-crawl/internal/vocabulary"
	"log"
	"runtime"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var source = "https://japanesetest4you.com/jlpt-n1-vocabulary-list/"

	runtime.GOMAXPROCS(4)

	uri := "mongodb://root:example@localhost:27017"

	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Connection error:", err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Ping error:", err)
	}

	fmt.Println("✅ Connected to MongoDB!")

	// Access a specific collection (example: "test" database, "users" collection)
	collection := client.Database("test").Collection("users")
	fmt.Printf("Collection: %v\n", collection.Name())

	repo, _ := vocabulary.NewVocalbularyRepository(client)
	vc := vocabulary.NewVocalbularyUseCaseImpl(repo)
	vc.Crawl(source)
}
