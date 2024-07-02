package main

import (
	"context"
	"go-msg/backend"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var Pool *pgxpool.Pool

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	pool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Unable to connect to db")
		os.Exit(1)
	}
	Pool = pool
	

	defer pool.Close()

	backend.Start(pool)
}
