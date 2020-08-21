package storage

import (
  "context"
  "os"
  "fmt"
  "github.com/jackc/pgx/v4/pgxpool"
)

func Start() (*pgxpool.Pool)  {
  dbURL := os.Getenv("DATABASE_URL")
  if dbURL == "" {
    dbURL = "postgresql://tinylibrary:password@localhost:5432/tinylibrary"
  }
  pool, err := pgxpool.Connect(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
  fmt.Println("Connected to the database", dbURL)
  return pool
}
