package api

import (
	"encoding/json"
	"go-msg/backend/internal/db"
	"log"
	"net/http"
)

type request struct {
	Sender string `json:"sender"`
}

func LoadMessages(w http.ResponseWriter, r *http.Request) {

	var req request
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

	log.Println(req.Sender)
	messages, err := db.GetMessages(req.Sender, user_id)
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
