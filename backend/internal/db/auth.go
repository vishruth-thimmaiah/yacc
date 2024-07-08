package db

import (
	"context"
	"errors"
	"yacc/backend/internal/helpers"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Setup(pool *pgxpool.Pool) {
	Pool = pool
}

func Login(email string, passwd string) (string, error) {

	var passwdHash, id string

	err := Pool.QueryRow(context.Background(), `select passwd,id from users where email = $1 limit 1`, email).Scan(&passwdHash, &id)
	if err != nil {
		return "", err
	}

	ok := helpers.VerifyHash(passwd, passwdHash)
	if !ok {
		return "", errors.New("Password is invalid")
	}

	session_id := uuid.New()

	_, err = Pool.Exec(context.Background(), `insert into session(sessionid, id) values($1, $2)`, session_id, id)
	if err != nil {
		return "", err
	}

	return session_id.String(), err
}

func Signup(email string, passwd string) (string, error) {
	_, err := Pool.Exec(context.Background(), `insert into users(email, passwd) values($1, $2)`, email, passwd)
	if err != nil {
		println(err)
		return "", err
	}

	session_id := uuid.New()

	_, err = Pool.Exec(context.Background(), `insert into session(sessionid, id) values($1, (select id from users where email = $2))`, session_id, email)
	if err != nil {
		return "", err
	}

	return session_id.String(), err
}

func SessionInfo(session_id string) (string, error) {
	var id string
	err := Pool.QueryRow(context.Background(), `select id from session where sessionid = $1 limit 1`, session_id).Scan(&id)
	if err != nil {
		return "", err
	}

	return id, err
}

func Logout(session_id string) error {
	_, err := Pool.Exec(context.Background(), `delete from session where sessionid = $1`, session_id)
	return err
} 
