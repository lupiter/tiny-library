package main

import (
    // "fmt"
    "log"
    "net/http"
    "github.com/lupiter/tiny-library/backend/internal/controllers"
    "github.com/lupiter/tiny-library/backend/internal/storage"
)

func main() {
    conn := storage.Start()
    defer conn.Close()
    storage := controllers.Storage{
      Pool: conn,
    }
    storage.Add()
    log.Fatal(http.ListenAndServe(":8080", nil))
}
