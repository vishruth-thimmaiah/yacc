package user

import (
	"encoding/json"
	"net/http"
	"yacc/backend/internal/db"
)

type NewUsername struct {
	Username string `json:"username"`
}

func ChangeUsername(w http.ResponseWriter, r *http.Request) {

	var username NewUsername
	json.NewDecoder(r.Body).Decode(&username)

	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user_id, err := db.UserInfo(session.Value)
	if err != nil {
		http.Error(w, "internal	error", http.StatusInternalServerError)
		return
	}

	db.ChangeUsername(username.Username, user_id)
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user_id, err := db.UserInfo(session.Value)
	if err != nil {
		http.Error(w, "internal	error", http.StatusInternalServerError)
		return
	}

	db.DeleteAccount(user_id)
}
