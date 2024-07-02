package db

import (
	"context"
	"log"
)

type Message struct {
	Message string `json:"message"`
	Sent    bool   `json:"sent"`
}

func GetMessages(sender string, id string) ([]Message, error) {
	query := "(select message, true from messages where senderid = $1 and recieverid = $2) union (select message, false from messages where senderid = $2 and recieverid = $1)"
	sent, serr := Pool.Query(context.Background(), query, sender, id)

	if serr != nil {
		return nil, serr
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

	log.Println(messages)
	return messages, nil
}
