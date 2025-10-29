package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChangeEvent struct {
	ID            bson.M `bson:"_id"`
	OperationType string `bson:"operationType"`
	FullDocument  bson.M `bson:"fullDocument"`
	Ns            bson.M `bson:"ns"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Connect to MongoDB (replica set required for change streams)
	uri := "mongodb://lucy:password@localhost:27017,localhost:27018,localhost:27019/?replicaSet=lucy-mongo&authSource=admin"

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	dbName := "inventory"
	collName := "items"
	collection := client.Database(dbName).Collection(collName)

	fmt.Printf("✅ Connected to MongoDB on: %s\n", uri)
	fmt.Printf("👂 Listening for inserts on: %s.%s\n", dbName, collName)
	fmt.Println("--------------------------------------------------")

	// 2. Simulate inserting a document in background
	go func() {
		time.Sleep(5 * time.Second) // wait until change stream starts
		doc := bson.M{
			"name":     "Sample Item",
			"category": "Testing",
			"price":    12.99,
			"created":  time.Now(),
		}
		res, err := collection.InsertOne(context.Background(), doc)
		if err != nil {
			log.Printf("❌ Failed to insert test document: %v", err)
			return
		}
		fmt.Printf("🧩 Inserted test document with ID: %v\n", res.InsertedID)
	}()

	// 3. Open change stream for insert operations
	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.D{
			{Key: "operationType", Value: "insert"},
		}}},
	}
	streamOptions := options.ChangeStream().SetFullDocument(options.UpdateLookup)

	cs, err := collection.Watch(context.Background(), pipeline, streamOptions)
	if err != nil {
		log.Fatal("Failed to open change stream:", err)
	}
	defer cs.Close(context.Background())

	// 4. Process new events
	for cs.Next(context.Background()) {
		var event ChangeEvent
		if err := cs.Decode(&event); err != nil {
			log.Printf("Error decoding event: %v", err)
			continue
		}

		fmt.Printf(">>> NEW DOCUMENT INSERTED (Type: %s)\n", event.OperationType)
		fmt.Printf("    - Collection: %s\n", event.Ns["coll"])
		fmt.Printf("    - Document ID: %v\n", event.FullDocument["_id"])
		fmt.Printf("    - Full Document Data: %v\n", event.FullDocument)
		fmt.Println("--------------------------------------------------")
	}

	if err := cs.Err(); err != nil {
		log.Println("Change Stream Error:", err)
	}
}
