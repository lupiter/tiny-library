package storage

import (
  "context"
  "fmt"
  "os"
  "github.com/lupiter/tiny-library/backend/internal/models"
  "github.com/jackc/pgx/v4/pgxpool"
)


func AllBooks(pool *pgxpool.Pool) ([]models.Book)  {
	var books []models.Book
	rows, err := pool.Query(context.Background(), "SELECT id, title, author, isbn, year, publisher, tags, max_loan_days, location, format FROM books;")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Query failed: %v\n", err)
		os.Exit(1)
	}

  for rows.Next() {
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
    rows.Scan(&id, &title, &author, &isbn, &year, &publisher, &tags, &max, &location, &format)
    book := models.Book{
        Identifier: id,
        Title: title,
        Author: author,
        ISBN: isbn,
        Year: year,
        Publisher: publisher,
        Tags: tags,
        MaxLoanDays: max,
        Location: location,
        Format: format,
    }
    books = append(books, book)
  }

  return books
}
