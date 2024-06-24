package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    MongoDBURI      string
    MongoDBDatabase string
}

func LoadConfig() Config {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Println("No .env file found")
    }

    return Config{
        MongoDBURI:      getEnv("MONGODB_URI", "mongodb://localhost:27017"),
        MongoDBDatabase: getEnv("MONGODB_DATABASE", "mydatabase"),
    }
}

func getEnv(key, defaultValue string) string {
    value := os.Getenv(key)
    if value == "" {
        return defaultValue
    }
    return value
}

