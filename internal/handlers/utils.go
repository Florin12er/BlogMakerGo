package handlers

import (
    "path/filepath"
    "runtime"
)

// getTemplatePath constructs the absolute path to the template file
func getTemplatePath(filename string) (string, error) {
    _, b, _, _ := runtime.Caller(0)
    basePath := filepath.Join(filepath.Dir(b), "../..")
    return filepath.Join(basePath, "web", "templates", filename), nil
}

