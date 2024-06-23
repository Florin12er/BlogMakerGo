package routes

import (
    "github.com/gorilla/mux"
    "Go/internal/handlers"
)

func SetupRoutes() *mux.Router {
    router := mux.NewRouter()

    router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
    router.HandleFunc("/about", handlers.AboutHandler).Methods("GET")

    return router
}

