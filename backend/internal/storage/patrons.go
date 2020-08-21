package storage

import (
  "context"
  "fmt"
  "os"
  "github.com/lupiter/tiny-library/backend/internal/models"
  "github.com/jackc/pgx/v4/pgxpool"
  "github.com/jackc/pgx/v4"
)

func PatronById(pool *pgxpool.Pool, patronId string) (models.Patron) {
  row := pool.QueryRow(context.Background(), "SELECT id, card_number, name, dob, address, active, max_loan_days FROM patrons WHERE id=$1;", patronId)
  patron := scanPatron(row)
  return patron
}

func scanPatron(row pgx.Row) (models.Patron) {
  var id string
  var cardNumber string
  var name string
  var dob string
  var address string
  var active bool
  var max int
  row.Scan(&id, &cardNumber, &name, &dob, &address, &active, &max)
  patron := models.Patron{
      Identifier: id,
      CardNumber: cardNumber,
      Name: name,
      DateOfBirth: dob,
      Address: address,
      Active: active,
      MaxLoanDays: max,
  }
  return patron
}

func AllPatrons(pool *pgxpool.Pool) ([]models.Patron)  {
	var patrons []models.Patron
	rows, err := pool.Query(context.Background(), "SELECT id, card_number, name, dob, address, active, max_loan_days FROM patrons;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

  for rows.Next() {
    patron := scanPatron(rows)
    patrons = append(patrons, patron)
  }

  return patrons
}
