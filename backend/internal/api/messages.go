package api

import (
	"encoding/json"
	"log"
	"net/http"
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
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Messages(w http.ResponseWriter, r *http.Request) {

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
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}

		var msg SendMessageRequest
		json.Unmarshal(message, &msg)

		saveMessage(msg.ChatId, user_id, msg.Message)

		err = c.WriteMessage(mt, []byte("etst"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func saveMessage(chat_id string, user_id string, message string) error {
	err := db.SendMessage(message, user_id, chat_id)
	return err

}
