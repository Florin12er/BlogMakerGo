package main

import (
    "log"
    "net/http"

    "Go/internal/config"
    "Go/internal/db"
    "Go/internal/routes"
)


func main() {
    cfg := config.LoadConfig()
    database.ConnectDB(cfg)

    router := routes.SetupRoutes()

    log.Println("Starting server on :8080")
    if err := http.ListenAndServe(":8080", router); err != nil {
        log.Fatalf("Could not start server: %s\n", err.Error())
    }
}

