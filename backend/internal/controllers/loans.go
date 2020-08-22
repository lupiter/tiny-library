package controllers

import (
	"fmt"
	// "log"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/lupiter/tiny-library/backend/internal/storage"
	"net/http"
)

func (s *Storage) AddLoans(r *mux.Router, prefix string) {
	r.HandleFunc(fmt.Sprintf("%s/loans", prefix), s.listLoans)
	r.HandleFunc(fmt.Sprintf("%s/loans/", prefix), s.listLoans)
	r.HandleFunc(fmt.Sprintf("%s/loans/{id}", prefix), s.oneLoan)
}

func (s *Storage) oneLoan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	loanId := vars["id"]
	loan := storage.LoanById(s.Pool, loanId)
	js, err := json.Marshal(loan)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}

func (s *Storage) listLoans(w http.ResponseWriter, r *http.Request) {
	all := storage.AllLoans(s.Pool)
	js, err := json.Marshal(all)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
}
