package db

import (
	"context"
)

type Message struct {
	Message string `json:"message"`
	Sent    bool   `json:"sent"`
}

func GetMessages(sender string, id string) ([]Message, error) {
	query := "(select message, true from messages where senderid = $1 and receiverid = $2) union (select message, false from messages where senderid = $2 and receiverid = $1)"
	sent, err := Pool.Query(context.Background(), query, sender, id)

	if err != nil {
		return nil, err
	}

	var messages []Message

	for sent.Next() {
		var message Message
		err := sent.Scan(&message.Message, &message.Sent)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func SendMessage(message string, id string, receiver string) error {
	_, err := Pool.Exec(context.Background(), `insert into messages(senderid, receiverid, message) values($1, $2, $3)`, id, receiver, message)
	return err
}
