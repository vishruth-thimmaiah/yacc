package messages

import (
	"encoding/json"
	"net/http"
	"time"
	"yacc/backend/internal/db"

	"github.com/gorilla/websocket"
)

type ShowMessagesRequest struct {
	Chat_id string `json:"chat_id"`
}

func LoadMessages(w http.ResponseWriter, r *http.Request) {

	var req ShowMessagesRequest
	json.NewDecoder(r.Body).Decode(&req)

	w.Header().Add("Content-type", "application/json")

	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user_id, err := db.SessionInfo(session.Value)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	messages, err := db.GetMessages(req.Chat_id, user_id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&messages)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

}

type SendMessageRequest struct {
	Message string `json:"message"`
	ChatId  string `json:"chat_id"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// TODO to be removed
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Chat_id     string    `json:"chat_id"`
	Receiver_id string    `json:"sender_id"`
	Date        time.Time `json:"date"`
	Message     string    `json:"message"`
}

func Messages(hub *Hub, w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user_id, err := db.SessionInfo(session.Value)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := Client{user_id: user_id, conn: c, hub: hub, sent: make(chan Message)}

	hub.register <- &client

	go client.read()
	go client.write()
}
