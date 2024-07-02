package db

import (
	"context"
	"errors"
	"go-msg/backend/internal/helpers"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Setup(pool *pgxpool.Pool) {
	Pool = pool
}

func Login(email string, passwd string) (string, error) {

	var passwdHash string
	var session_id string

	err := Pool.QueryRow(context.Background(), `select passwd, sessionid from users where email = $1 limit 1`, email).Scan(&passwdHash, &session_id)

	if err != nil {
		log.Println(err)
		return "", err
	}

	log.Println(passwdHash)
	ok := helpers.VerifyHash(passwd, passwdHash)

	if !ok {
		return "", errors.New("Password is invalid")
	}

	return session_id, err
}

func Signup(email string, passwd string) error {
	_, err := Pool.Exec(context.Background(), `insert into users(email, passwd) values($1, $2)`, email, passwd)
	return err
}
