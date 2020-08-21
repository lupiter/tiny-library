package main

import (
	"github.com/gorilla/mux"
	"github.com/lupiter/tiny-library/backend/internal/controllers"
	"github.com/lupiter/tiny-library/backend/internal/storage"
	"log"
	"net/http"
	"os"
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
	s.AddRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
