package api

import (
	"encoding/json"
	"go-msg/backend/internal/db"
	"net/http"
)

func Contacts(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-type", "application/json")

	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user_id, err := db.SessionInfo(session.Value)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	contacts, err := db.GetContacts(user_id)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(&contacts)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

}
