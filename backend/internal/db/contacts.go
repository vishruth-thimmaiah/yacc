package db

import (
	"context"
	"log"
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

func AddContact(id string, new_email string) error {
	_, err := Pool.Exec(context.Background(), "insert into contacts(user1, user2) values($1, (select id from users where email=$2))", id, new_email)
	return err
}

func GetReceipient(chat_id string, sender_id string) string {
	var query = `select case
			when user1 = $1 then user2
			when user2 = $1 then user1
		end
		from contacts
		where chat_id =$2
		limit 1`

	var receipient string
	err := Pool.QueryRow(context.Background(), query, sender_id, chat_id).Scan(&receipient)
	if err != nil {
		log.Println(err)
		return ""
	}
	return receipient
}
