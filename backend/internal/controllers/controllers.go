package controllers

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func (s *Storage) AddRoutes(router *mux.Router, prefix string) {
	s.AddBooks(router, prefix)
	s.AddPatrons(router, prefix)
	s.AddLoans(router, prefix)
}

type Storage struct {
	Pool *pgxpool.Pool
}
