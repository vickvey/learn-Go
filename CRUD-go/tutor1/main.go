package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title  string
	Author string
}

func main() {
	fmt.Println("Hello World!!")

	// Specify a server URI to connect to
	uri := "mongodb://localhost:27017"

	// Specify the Stable API version in the ClientOptions object
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	// Pass in the URI and the ClientOptions to the Client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI))
	if err != nil {
		panic(err)
	}

	// using the connection to insert data to database
	coll := client.Database("db").Collection("books")
	doc := Book{Title: "Atonement", Author: "Ian McEwan"}

	result, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)

	// insert many documents to database
	coll2 := client.Database("db").Collection("favorite_books")
	docs := []interface{}{
		Book{Title: "Cat's Cradle", Author: "Kurt Vonnegut Jr."},
		Book{Title: "In Memory of Memory", Author: "Maria Stepanova"},
		Book{Title: "Pride and Prejudice", Author: "Jane Austen"},
	}

	result2, err := coll2.InsertMany(context.TODO(), docs)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Documents inserted: %v\n", len(result2.InsertedIDs))

	for _, id := range result2.InsertedIDs {
		fmt.Printf("Inserted document with _id: %v\n", id)
	}
}
