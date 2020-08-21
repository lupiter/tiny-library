package storage

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/lupiter/tiny-library/backend/internal/models"
	"os"
)

const bookColumns = "id, title, author, isbn, year, publisher, tags, max_loan_days, location, format"

func BookById(pool *pgxpool.Pool, bookId string) models.Book {
	row := pool.QueryRow(context.Background(), "SELECT "+bookColumns+" FROM books WHERE id=$1;", bookId)
	book := scanBook(row)
	return book
}

func AddBook(pool *pgxpool.Pool, book models.Book) {
	_, err := pool.Exec(
		context.Background(),
		"INSERT INTO books ("+bookColumns+") VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		book.Identifier,
		book.Title,
		book.Author,
		book.ISBN,
		book.Year,
		book.Publisher,
		book.Tags,
		book.MaxLoanDays,
		book.Location,
		book.Format,
	)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Exec failed: %v\n", err)
		os.Exit(1)
	}
}

func scanBook(row pgx.Row) models.Book {
	var id string
	var title string
	var author string
	var isbn string
	var year string
	var publisher string
	var tags []string
	var max int
	var location string
	var format string
	row.Scan(&id, &title, &author, &isbn, &year, &publisher, &tags, &max, &location, &format)
	book := models.Book{
		Identifier:  id,
		Title:       title,
		Author:      author,
		ISBN:        isbn,
		Year:        year,
		Publisher:   publisher,
		Tags:        tags,
		MaxLoanDays: max,
		Location:    location,
		Format:      format,
	}
	return book
}

func AllBooks(pool *pgxpool.Pool) []models.Book {
	var books []models.Book
	rows, err := pool.Query(context.Background(), "SELECT id, title, author, isbn, year, publisher, tags, max_loan_days, location, format FROM books;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

	for rows.Next() {
		book := scanBook(rows)
		books = append(books, book)
	}

	return books
}
