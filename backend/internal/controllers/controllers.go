package controllers

import (
  "github.com/jackc/pgx/v4/pgxpool"
)

func (s *Storage) Add() {
  s.AddBooks("")
  s.AddPatrons("")
  s.AddLoans("")
}


type Storage struct {
  Pool *pgxpool.Pool
}
