FROM golang:1.14 as build-stage
WORKDIR /app
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o tiny ./cmd/tinylibraryservice

FROM alpine:latest
COPY --from=build-stage /app /app
EXPOSE 8080
CMD ["./tiny"]