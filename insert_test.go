package main_test

import (
	"context"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestInsertDocument(t *testing.T) {
	// Define the MongoDB connection URI
	uri := "mongodb://mongo-primary:27017,mongo-secondary1:27017,mongo-secondary2:27017/?replicaSet=rs0"

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connect to the MongoDB replica set
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		t.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			t.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()

	// Select the database and collection
	collection := client.Database("inventory").Collection("items")

	// Define the document to insert
	document := bson.D{
		{Key: "name", Value: "item1"},
		{Key: "quantity", Value: 100},
		{Key: "price", Value: 10.99},
	}

	// Insert the document
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		t.Fatalf("Failed to insert document: %v", err)
	}

	// Verify the insertion
	if result.InsertedID == nil {
		t.Fatalf("Document insertion failed, no ID returned")
	}

	t.Logf("Document inserted successfully with ID: %v", result.InsertedID)
}
