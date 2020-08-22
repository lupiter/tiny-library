package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lupiter/tiny-library/backend/internal/controllers"
	"github.com/lupiter/tiny-library/backend/internal/storage"
)

func main() {
	pool := storage.Start()
	defer pool.Close()
	s := controllers.Storage{
		Pool: pool,
	}
	storage.LoadBooks(pool, os.Getenv("DATA_BOOKS"))
	storage.LoadPatrons(pool, os.Getenv("DATA_PATRONS"))
	storage.LoadLoans(pool, os.Getenv("DATA_LOANS"))

	r := mux.NewRouter()
	s.AddRoutes(r, "/api/v0")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
