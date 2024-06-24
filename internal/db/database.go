package database

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "Go/internal/config"
)

var Client *mongo.Client

func ConnectDB(cfg config.Config) {
    clientOptions := options.Client().ApplyURI(cfg.MongoDBURI)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB")
    Client = client
}

func GetCollection(collectionName string) *mongo.Collection {
    cfg := config.LoadConfig()
    return Client.Database(cfg.MongoDBDatabase).Collection(collectionName)
}
