package messages

import (
	"log"
)

type Hub struct {
	clients    []*Client
	message    chan Message
	register   chan *Client
	unregister chan *Client
}

func CreateHub() *Hub {
	return &Hub{
		clients:    nil,
		message:    make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients = append(h.clients, client)
			log.Println("new client registered")

		case client := <-h.unregister:
			for client_index := range h.clients {
				if h.clients[client_index] == client {
					h.clients = append(h.clients[:client_index], h.clients[client_index+1:]...)
					break
				}
			}
			log.Println("client unregistered")

		case message := <-h.message:
			for _, client := range h.clients {
				if client.user_id == message.Receiver_id {
					go client.Write(message)
					break
				}
			}
		}
	}
}
