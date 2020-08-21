package storage

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lupiter/tiny-library/backend/internal/models"
	"io/ioutil"
	"os"
)

const patronColumns = "card_number, name, dob, address, active, max_loan_days"

func PatronById(pool *pgxpool.Pool, patronId string) models.Patron {
	row := pool.QueryRow(context.Background(), "SELECT id, "+patronColumns+" FROM patrons WHERE id=$1;", patronId)
	patron := scanPatron(row)
	return patron
}

func LoadPatrons(pool *pgxpool.Pool) {
	patronDataFile := os.Getenv("DATA_PATRONS")
	if patronDataFile != "" {
		patronData, err := ioutil.ReadFile(patronDataFile)
		if err != nil {
			// TODO
		}
		var patrons []models.Patron
		err = json.Unmarshal(patronData, &patrons)
		if err != nil {
			// TODO
		}
		for _, patron := range patrons {
			AddPatron(pool, patron)
		}
	}
}

func scanPatron(row pgx.Row) models.Patron {
	var id string
	var cardNumber string
	var name string
	var dob string
	var address string
	var active bool
	var max int
	row.Scan(&id, &cardNumber, &name, &dob, &address, &active, &max)
	patron := models.Patron{
		Identifier:  id,
		CardNumber:  cardNumber,
		Name:        name,
		DateOfBirth: dob,
		Address:     address,
		Active:      active,
		MaxLoanDays: max,
	}
	return patron
}

func AddPatron(pool *pgxpool.Pool, patron models.Patron) models.Patron {
	var id string
	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO patrons ("+patronColumns+") VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		patron.Identifier,
		patron.CardNumber,
		patron.Name,
		patron.DateOfBirth,
		patron.Address,
		patron.Active,
		patron.MaxLoanDays,
	).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
	patron.Identifier = id
	return patron
}

func AllPatrons(pool *pgxpool.Pool) []models.Patron {
	var patrons []models.Patron
	rows, err := pool.Query(context.Background(), "SELECT id, "+patronColumns+" FROM patrons;")
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
