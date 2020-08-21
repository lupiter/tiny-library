package main

import (
    // "fmt"
    "log"
    "net/http"
    "github.com/lupiter/tiny-library/backend/internal/controllers"
    "github.com/lupiter/tiny-library/backend/internal/storage"
    "github.com/gorilla/mux"
)

func main() {
    conn := storage.Start()
    defer conn.Close()
    storage := controllers.Storage{
      Pool: conn,
    }
    r := mux.NewRouter()
    storage.AddRoutes(r)
    http.Handle("/", r)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
