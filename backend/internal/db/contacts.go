package db

import (
	"context"
)

type Contact struct {
	ChatId string `json:"chat_id"`
	Name   string `json:"name"`
}

func GetContacts(id string) ([]Contact, error) {
	var query = `
		SELECT c.chat_id, u.username FROM contacts AS c
		INNER JOIN users AS u ON (
			c.user1 = $1
			AND c.user2 = u.id
		)
		OR (
		c.user2 = $1
			AND c.user1 = u.id
		)`

	rows, err := Pool.Query(context.Background(), query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.ChatId, &contact.Name)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func AddContact(id string, new_id string) error {
	_, err := Pool.Exec(context.Background(), "insert into contacts(user1, user2) values($1, $2)", id, new_id)
	return err
}
