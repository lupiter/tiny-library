package controllers

import (
    "fmt"
    // "log"
    "net/http"
    "encoding/json"
    "github.com/lupiter/tiny-library/backend/internal/storage"
    "github.com/gorilla/mux"
)

func (s *Storage) AddBooks(r *mux.Router, prefix string) {
    r.HandleFunc(fmt.Sprintf("%s/books", prefix), s.listBooks)
    r.HandleFunc(fmt.Sprintf("%s/books/", prefix), s.listBooks)
    r.HandleFunc(fmt.Sprintf("%s/books/{id}", prefix), s.oneBook)
}

func (s *Storage) oneBook(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  bookId := vars["id"]
  book := storage.BookById(s.Pool, bookId)
  js, err := json.Marshal(book)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.Write(js)
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
