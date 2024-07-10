package main

import (
	"context"
	"log"
	"os"
	"yacc/backend"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func init() {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = "production"
	}

	log.Printf("running in %s mode\n", env)
	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		log.Println("Error loading .env")
	}
}

func main() {

	pool, err := pgxpool.New(context.Background(), os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal("Unable to connect to db")
	}

	defer pool.Close()

	backend.Start(pool)
}
