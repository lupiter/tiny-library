package controllers

import (
    "fmt"
    // "log"
    "net/http"
)

func (s *Storage) AddLoans(prefix string) {
    http.HandleFunc(fmt.Sprintf("%s/loans", prefix), s.listLoans)
}

func (s *Storage) listLoans(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
