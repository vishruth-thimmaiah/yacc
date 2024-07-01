package main

import (
	"context"
	"go-msg/backend"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Test string = "test"
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Unable to connect to db")
		os.Exit(1)
	}
	defer conn.Close()

	backend.Start(conn)
}
