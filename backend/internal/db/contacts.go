package db

import (
	"context"
)

type Contact struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

func GetContacts(id string) ([]Contact, error) {
	rows, err := Pool.Query(context.Background(), "select id, username from users where id in (select senderid from contacts where receiverid = $1)", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.Id, &contact.Name)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
