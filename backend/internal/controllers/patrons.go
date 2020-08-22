package controllers

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lupiter/tiny-library/backend/internal/storage"
	"net/http"
)

func (s *Storage) AddPatrons(r *mux.Router, prefix string) {
	r.HandleFunc(fmt.Sprintf("%s/people", prefix), s.listPatrons)
	r.HandleFunc(fmt.Sprintf("%s/people/", prefix), s.listPatrons)
	r.HandleFunc(fmt.Sprintf("%s/people/{id}", prefix), s.onePatron)
	r.HandleFunc(fmt.Sprintf("%s/people/{id}/loans", prefix), s.loansForPatron)
}

func (s *Storage) onePatron(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patronId := vars["id"]
	patron := storage.PatronById(s.Pool, patronId)
	js, err := json.Marshal(patron)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Storage) listPatrons(w http.ResponseWriter, r *http.Request) {
	all := storage.AllPatrons(s.Pool)
	js, err := json.Marshal(all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (s *Storage) loansForPatron(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patronId := vars["id"]
	loans := storage.LoansByPatron(s.Pool, patronId)
	js, err := json.Marshal(loans)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
