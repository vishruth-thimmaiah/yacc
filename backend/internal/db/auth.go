package db

import (
	"context"
	"errors"
	"yacc/backend/internal/helpers"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Setup(pool *pgxpool.Pool) {
	Pool = pool
}

func Login(email string, passwd string) (string, error) {

	var hashed_passwd, user_id string

	err := Pool.QueryRow(context.Background(), `select passwd,id from users where email = $1 limit 1`, email).Scan(&hashed_passwd, &user_id)
	if err != nil {
		return "", err
	}

	ok := helpers.VerifyHash(passwd, hashed_passwd)
	if !ok {
		return "", errors.New("Password is invalid")
	}

	session_id := uuid.New()

	_, err = Pool.Exec(context.Background(), `insert into session(sessionid, id) values($1, $2)`, session_id, user_id)
	if err != nil {
		return "", err
	}

	return session_id.String(), err
}

func Signup(email string, passwd string, username string) (string, *pgconn.PgError) {
	_, err := Pool.Exec(context.Background(), `insert into users(email, passwd, username) values($1, $2, $3)`, email, passwd, username)
	if err != nil {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		return "", pgErr
	}

	session_id := uuid.New()

	_, err = Pool.Exec(context.Background(), `insert into session(sessionid, id) values($1, (select id from users where email = $2))`, session_id, email)
	if err != nil {
		var pgErr *pgconn.PgError
		errors.As(err, &pgErr)
		return "", pgErr
	}

	return session_id.String(), nil
}

func UserInfo(session_id string) (string, error) {
	var user_id string
	err := Pool.QueryRow(context.Background(), `select id from session where sessionid = $1 limit 1`, session_id).Scan(&user_id)
	if err != nil {
		return "", err
	}

	return user_id, err
}

func Logout(session_id string) error {
	_, err := Pool.Exec(context.Background(), `delete from session where sessionid = $1`, session_id)
	return err
}

func ChangeUsername(username string, user_id string) error {
	_, err := Pool.Exec(context.Background(), `update users set username=$1 where id=$2`, username, user_id)
	return err
}

func DeleteAccount(user_id string) error {
	_, err := Pool.Exec(context.Background(), `delete from users where id=$1`, user_id)
	return err
}
