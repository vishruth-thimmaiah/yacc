package api

import (
	"encoding/json"
	"net/http"
	"yacc/backend/internal/db"
)

type AddContactReq struct {
	Email string `json:"email"`
}

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

func NewContact(w http.ResponseWriter, r *http.Request) {

	var req AddContactReq
	json.NewDecoder(r.Body).Decode(&req)

	if req.Email == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

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

	err = db.AddContact(user_id, req.Email)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
}
