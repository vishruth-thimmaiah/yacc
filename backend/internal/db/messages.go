package db

import (
	"context"
	"time"
)

type Message struct {
	Message string `json:"message"`
	Date    time.Time `json:"date"`
	Sent    bool   `json:"sent"`
}

func GetMessages(chat_id string, sender_id string) ([]Message, error) {
	sent, err := Pool.Query(context.Background(), "select message, date, senderid!=$1 from messages where chat_id = $2", sender_id, chat_id)

	if err != nil {
		return nil, err
	}

	var messages []Message

	for sent.Next() {
		var message Message
		err := sent.Scan(&message.Message, &message.Date, &message.Sent)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func SendMessage(message string, sender_id string, chat_id string) error {
	_, err := Pool.Exec(context.Background(), `insert into messages(senderid, chat_id, message) values($1, $2, $3)`, sender_id, chat_id, message)
	return err
}
