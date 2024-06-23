package handlers

import (
    "html/template"
    "net/http"
    "log"
)

// AboutHandler handles the about page requests
func AboutHandler(w http.ResponseWriter, r *http.Request) {
    tmplPath, err := getTemplatePath("about.html")
    if err != nil {
        log.Println("Error getting template path:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    layoutPath, err := getTemplatePath("layout.html")
    if err != nil {
        log.Println("Error getting layout path:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    tmpl, err := template.ParseFiles(layoutPath, tmplPath)
    if err != nil {
        log.Println("Error parsing templates:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    data := struct {
        Title       string
        Description string
    }{
        Title:       "About Us",
        Description: "This is the about page of our Go web application.",
    }

    log.Println("Rendering about page")
    err = tmpl.ExecuteTemplate(w, "layout.html", data)
    if err != nil {
        log.Println("Error executing template:", err)
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

