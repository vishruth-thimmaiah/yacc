package messages

import (
	"encoding/json"
	"log"
	"yacc/backend/internal/db"

	"github.com/gorilla/websocket"
)

type Client struct {
	hub     *Hub
	user_id string
	conn    *websocket.Conn
}

func (c *Client) read() {

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, response, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var message Message
		json.Unmarshal(response, &message)
		message.Receiver_id = db.GetReceipient(message.Chat_id, c.user_id)

		c.hub.message <- message

		if message.Type == "message" {
			db.SaveMessage(message.Message, c.user_id, message.Chat_id, message.Attachment_url)
		}
	}
}

func (c *Client) Write(message Message) {

	response, err := json.Marshal(message)
	if err != nil {
		return
	}
	err = c.conn.WriteMessage(websocket.TextMessage, response)
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			log.Printf("error: %v", err)
		}
		return
	}
}
