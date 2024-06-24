package handlers

import (
    "encoding/json"
    "github.com/gorilla/mux"
    "math/rand"
    "net/http"
    "strconv"
)

type Blog struct {
    ID      string  `json:"id"`
    Title   string  `json:"title"`
    Links   string  `json:"links"`
    Content string  `json:"content"`
    Author  *Author `json:"author"`
}

type Author struct {
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}

var blogs []Blog

func InitBlogs() {
    blogs = append(blogs, Blog{
        ID:      "1",
        Title:   "Blog",
        Links:   "best value, computer",
        Content: "this is my blog",
        Author: &Author{
            Firstname: "John",
            Lastname:  "Doe",
        },
    })

    blogs = append(blogs, Blog{
        ID:      "2",
        Title:   "Blog2",
        Links:   "best value, cayomputer",
        Content: "this is not my blog",
        Author: &Author{
            Firstname: "Johnu",
            Lastname:  "Moe",
        },
    })
}

func GetBlogs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(blogs)
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range blogs {
        if item.ID == params["id"] {
            blogs = append(blogs[:index], blogs[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(blogs)
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range blogs {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    // If blog not found, return appropriate response (optional)
    http.NotFound(w, r)
}

func CreateBlog(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var blog Blog
    _ = json.NewDecoder(r.Body).Decode(&blog)
    blog.ID = strconv.Itoa(rand.Intn(1000000000000000))
    blogs = append(blogs, blog)
    json.NewEncoder(w).Encode(blog)
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    var updatedBlog Blog
    _ = json.NewDecoder(r.Body).Decode(&updatedBlog)

    for index, item := range blogs {
        if item.ID == params["id"] {
            // Update the blog with new data
            blogs[index] = updatedBlog
            json.NewEncoder(w).Encode(updatedBlog)
            return
        }
    }

    // If blog not found, return appropriate response (optional)
    http.NotFound(w, r)
}

