package api

import (
	"encoding/json"
	"net/http"
	"yacc/backend/internal/db"
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
	Message  string `json:"message"`
	ChatId string `json:"chat_id"`
}

func SendMessage(w http.ResponseWriter, r *http.Request) {

	var req SendMessageRequest
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

	err = db.SendMessage(req.Message, user_id, req.ChatId)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
