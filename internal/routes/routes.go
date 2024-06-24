package routes

import (
	"Go/internal/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()
	handlers.InitBlogs()
	router.HandleFunc("/", handlers.HomeHandler).Methods("GET")
	router.HandleFunc("/about", handlers.AboutHandler).Methods("GET")
	router.HandleFunc("/example", handlers.ExampleHandler).Methods("GET")
    router.HandleFunc("/blogs", handlers.GetBlogs).Methods("GET")
    router.HandleFunc("/blogs/{id}", handlers.GetBlog).Methods("GET")
    router.HandleFunc("/blogs", handlers.CreateBlog).Methods("POST")
    router.HandleFunc("/blogs/{id}", handlers.UpdateBlog).Methods("PUT")
    router.HandleFunc("/blogs/{id}", handlers.DeleteBlog).Methods("DELETE")

	return router
}
