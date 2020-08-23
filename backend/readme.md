# Tiny Library Backend

## Running Locally

If you have the go toolchain already installed, you can build and run with `go run ./cmd/tinylibraryservice`

## Environment Variables

This service uses the following environment variables

Name         | Default    | Description
-------------|------------|-------------
DATA_BOOKS   | None       | If supplied, a JSON file containing books to load on startup
DATA_PATRONS | None       | If supplied, a JSON file containing patrons to load on startup
DATA_LOANS   | None       | If supplied, a JSON file containing loans to load on startup
DATABASE_URL | postgresql://tinylibrary:password@localhost:5432/tinylibrary | The URL for the database connection
