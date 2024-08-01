package messages

import (
	"encoding/json"
	"net/http"
	"time"
	"yacc/backend/internal/db"
	"yacc/backend/internal/helpers"
	"yacc/backend/internal/s3"

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

	user_id, err := db.UserInfo(session.Value)
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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// TODO to be removed
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Chat_id        string    `json:"chat_id"`
	Receiver_id    string    `json:"-"`
	Date           time.Time `json:"date"`
	Message        string    `json:"message"`
	Attachment_url string    `json:"attachment"`
}

type Attachment struct {
	Url string `json:"url"`
}

func Messages(hub *Hub, w http.ResponseWriter, r *http.Request) {

	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	user_id, err := db.UserInfo(session.Value)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	client := Client{user_id: user_id, conn: c, hub: hub}

	hub.register <- &client

	go client.read()
}

func Attachments(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "error reading file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content_type := header.Header.Get("Content-Type")
	ok := helpers.MatchFileType(content_type)
	if !ok {
		http.Error(w, "invalid file", http.StatusBadRequest)
		return
	}

	image_name := r.FormValue("image_name")

	var attachment Attachment

	attachment.Url, err = bucket.Upload(file, header, image_name, content_type)

	err = json.NewEncoder(w).Encode(&attachment)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
