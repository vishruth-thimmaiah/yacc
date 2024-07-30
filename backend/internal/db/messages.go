package db

import (
	"context"
	"time"
)

type Message struct {
	Message    string    `json:"message"`
	Date       time.Time `json:"date"`
	Sent       bool      `json:"sent"`
	Attachment string    `json:"attachment"`
}

func GetMessages(chat_id string, sender_id string) ([]Message, error) {
	sent, err := Pool.Query(context.Background(), "select message, date, senderid!=$1, attachment from messages where chat_id = $2", sender_id, chat_id)

	if err != nil {
		return nil, err
	}

	var messages []Message

	for sent.Next() {
		var message Message
		err := sent.Scan(&message.Message, &message.Date, &message.Sent, &message.Attachment)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}

func SaveMessage(message string, sender_id string, chat_id string, attachment string) error {
	query := `insert into messages(senderid, chat_id, message, attachment) values($1, $2, $3, $4)`
	_, err := Pool.Exec(context.Background(), query, sender_id, chat_id, message, attachment)
	return err
}
