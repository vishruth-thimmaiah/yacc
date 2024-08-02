package db

import "context"


func ChangeUsername(username string, user_id string) error {
	_, err := Pool.Exec(context.Background(), `update users set username=$1 where id=$2`, username, user_id)
	return err
}

func DeleteAccount(user_id string) error {
	_, err := Pool.Exec(context.Background(), `delete from users where id=$1`, user_id)
	return err
}
