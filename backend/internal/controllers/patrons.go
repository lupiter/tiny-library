package controllers

import (
    "fmt"
    // "log"
    "net/http"
)

func (s *Storage) AddPatrons(prefix string) {
    http.HandleFunc(fmt.Sprintf("%s/people", prefix), s.listPatrons)
}

func (s *Storage) listPatrons(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
