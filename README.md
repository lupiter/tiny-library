# tiny-library
 
## Building

You will need [Docker](https://www.docker.com/products/docker-desktop) to run this project. Once Docker is installed, simply run `docker-compose up` from this directory. The applicaiton will be available at http://localhost/library

## Building for Development

### Database

You will need [PostgreSQL](https://www.postgresql.org), on Mac the easiest way is with [Postgress.app](https://postgresapp.com).

In the `db` directory there is a shell script to create the appropriate database and user. Once created, you should be able to connect and run the sql in the same directory, which will create the tables.

### Backend

You will need a [Go](https://golang.org) environment. The easiest way on Mac is to install go via homebrew, but the official distributions are also fine.

You can start the backend with `go run ./cmd/tinylibraryservice` from the `backend` directory, and it will be available at http://localhost:8080/api/v0. For further details on available environment variables etc, please see the README.md in that directory.

### Frontend

You will need a [Node.js](https://nodejs.org/en/) environment. The official distribution or a homebrew or apt-get install are fine.

You can start the frontend with `npm run serve` from the `frontend/tiny-library` directory, and it will be available at http://localhost:8080/ (if the backend is already running, it will use port 8081). For further details please see the README.md in that directory.