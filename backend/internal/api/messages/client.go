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
	sent    chan Message
}

func (c *Client) read() {

	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, response, err := c.conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}

		var message Message
		json.Unmarshal(response, &message)
		message.Receiver_id = db.GetReceipient(message.Chat_id, c.user_id)

		c.hub.message <- message

		db.SaveMessage(message.Message, c.user_id, message.Chat_id)
	}
}

func (c *Client) write() {
	for {
		select {
		case message := <-c.sent:
			response, err := json.Marshal(message)
			if err != nil {
				continue
			}
			log.Println(string(response))
			err = c.conn.WriteMessage(websocket.TextMessage, response)
			if err != nil {
				log.Println(err)
				break
			}
		}
	}
}
