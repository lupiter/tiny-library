package controllers

import (
    "fmt"
    // "log"
    "net/http"
    "encoding/json"
    "github.com/lupiter/tiny-library/backend/internal/storage"
)

func (s *Storage) AddBooks(prefix string) {
    http.HandleFunc(fmt.Sprintf("%s/books", prefix), s.listBooks)
}

func (s *Storage) listBooks(w http.ResponseWriter, r *http.Request) {
    all := storage.AllBooks(s.Pool)
    js, err := json.Marshal(all)
    if err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
      return
    }

    w.Header().Set("Content-Type", "application/json")
    w.Write(js)
}
