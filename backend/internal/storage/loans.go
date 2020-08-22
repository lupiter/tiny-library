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

const loanSelect = `
  SELECT
    loans.id, loans.lent, loans.due_back, loans.returned,
    books.id, books.title, books.author, books.isbn, books.year, books.publisher, books.tags, books.max_loan_days, books.location, books.format,
    patrons.id, patrons.card_number, patrons.name, patrons.dob, patrons.address, patrons.active, patrons.max_loan_days
  FROM loans, patrons, books`

func LoanById(pool *pgxpool.Pool, loanId string) models.Loan {
	row := pool.QueryRow(context.Background(), loanSelect+" WHERE loans.id=$1 AND loans.patron = patrons.id AND loans.book = books.id;", loanId)
	loan := scanLoan(row)
	return loan
}

func AddLoan(pool *pgxpool.Pool, loan models.Loan) {
	var id int
	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO loans (book, patron, lent, due_back, returned) VALUES ($1, $2, $3, $4, $5) returning id;",
		loan.Book.Identifier,
		loan.Patron.Identifier,
		loan.Lent,
		loan.DueBack,
		loan.Returned,
	).Scan(&id)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n loan %v\n", err, loan)
		os.Exit(1)
	}
	loan.Identifier = id
}

func LoadLoans(pool *pgxpool.Pool, dataFile string) {
	if dataFile != "" {
		data, err := ioutil.ReadFile(dataFile)
		if err != nil {
			// TODO
		}
		var loans []models.Loan
		err = json.Unmarshal(data, &loans)
		if err != nil {
			// TODO
		}
		for _, loan := range loans {
			AddLoan(pool, loan)
		}
	}
}

func LoansByPatron(pool *pgxpool.Pool, patronId string) []models.Loan {
	var loans []models.Loan
	rows, err := pool.Query(context.Background(), loanSelect+" WHERE patron=$1 AND loans.patron = patrons.id AND loans.book = books.id;", patronId)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		loan := scanLoan(rows)
		loans = append(loans, loan)
	}

	return loans
}

func scanLoan(row pgx.Row) models.Loan {
	var id int
	var lent string
	var dueBack string
	var returned string

	var bookId int
	var title string
	var author string
	var isbn string
	var year string
	var publisher string
	var tags []string
	var maxBook int
	var location string
	var format string

	var patronId int
	var card string
	var name string
	var dob string
	var address string
	var active bool
	var maxPatron int
	row.Scan(&id, &lent, &dueBack, &returned, &bookId, &title, &author, &isbn, &year, &publisher, &tags, &maxBook, &location, &format, &patronId, &card, &name, &dob, &address, &active, &maxPatron)
	book := models.Book{
		Identifier:  bookId,
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		Year:        year,
		Publisher:   publisher,
		Tags:        tags,
		MaxLoanDays: maxBook,
		Location:    location,
		Format:      format,
	}
	patron := models.Patron{
		Identifier:  id,
		CardNumber:  card,
		Name:        name,
		DateOfBirth: dob,
		Address:     address,
		Active:      active,
		MaxLoanDays: maxPatron,
	}
	loan := models.Loan{
		Identifier: id,
		Book:       book,
		Patron:     patron,
		Lent:       lent,
		DueBack:    dueBack,
		Returned:   returned,
	}
	return loan
}

func AllLoans(pool *pgxpool.Pool) []models.Loan {
	var loans []models.Loan
	rows, err := pool.Query(context.Background(), loanSelect+" WHERE loans.patron = patrons.id AND loans.book = books.id;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		loan := scanLoan(rows)
		loans = append(loans, loan)
	}

	return loans
}
