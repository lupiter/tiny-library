package main

import (
	// "fmt"
	"github.com/gorilla/mux"
	"github.com/lupiter/tiny-library/backend/internal/controllers"
	"github.com/lupiter/tiny-library/backend/internal/storage"
	"log"
	"net/http"
)

func main() {
	pool := storage.Start()
	defer pool.Close()
	s := controllers.Storage{
		Pool: pool,
	}
	storage.LoadBooks(pool)
	storage.LoadPatrons(pool)
	storage.LoadLoans(pool)

	r := mux.NewRouter()
	s.AddRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
