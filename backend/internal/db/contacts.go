package db

import (
	"context"
)

type Contact struct {
	Name string `json:"name"`
}

func GetContacts(id string) ([]Contact, error) {
	rows, err := Pool.Query(context.Background(), "select username from users where id = (select senderid from contacts where recieverid = $1)", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []Contact

	for rows.Next() {
		var contact Contact
		err = rows.Scan(&contact.Name)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}
