package handlers

import (
    "context"
    "net/http"
    "log"
    "time"

    "Go/internal/db"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

// ExampleHandler demonstrates a basic MongoDB operation
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    collection := database.GetCollection("example")

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // Insert a document
    doc := bson.D{{"name", "John Doe"}, {"age", 30}}
    _, err := collection.InsertOne(ctx, doc)
    if err != nil {
        log.Println("Error inserting document:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Find a document
    var result bson.M
    err = collection.FindOne(ctx, bson.D{{"name", "John Doe"}}).Decode(&result)
    if err == mongo.ErrNoDocuments {
        http.Error(w, "No document found", http.StatusNotFound)
        return
    } else if err != nil {
        log.Println("Error finding document:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    log.Println("Found document:", result)
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Document inserted and found successfully"))
}

