package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func (s *Storage) AddRoutes(router *mux.Router) {
	s.AddBooks(router, "")
	s.AddPatrons(router, "")
	s.AddLoans(router, "")
}

type Storage struct {
	Pool *pgxpool.Pool
}
