package controllers

import (
  "github.com/jackc/pgx/v4/pgxpool"
  "github.com/gorilla/mux"
)

func (s *Storage) AddRoutes(router *mux.Router) {
  s.AddBooks(router, "")
  s.AddPatrons(router, "")
  s.AddLoans(router, "")
}


type Storage struct {
  Pool *pgxpool.Pool
}
