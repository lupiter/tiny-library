FROM golang:alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify
COPY . .
RUN go build -o tiny ./cmd/tinylibraryservice
EXPOSE 8080
CMD ["./tiny"]