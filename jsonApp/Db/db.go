package db

import (
	"context"
	"fmt"
	"jsonApp/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DBClient is a global variable for the MongoDB client
var DBClient *mongo.Client

func InitMongo() {
	_, err := ConnectDB(config.Mongo_URI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
}

// ConnectDB initializes a connection to MongoDB and returns the client
func ConnectDB(uri string) (*mongo.Client, error) {
	// Set client options
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		return nil, err
	}

	// Check the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	DBClient = client // Set the global client
	return client, nil
}

// GetDB returns the database instance
func GetDB(databaseName string) *mongo.Database {
	if DBClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	return DBClient.Database(databaseName)
}

// CloseDB closes the MongoDB client connection
func CloseDB() {
	if DBClient != nil {
		err := DBClient.Disconnect(context.Background())
		if err != nil {
			log.Printf("Failed to disconnect MongoDB: %v", err)
		} else {
			fmt.Println("Disconnected from MongoDB")
		}
	}
}
