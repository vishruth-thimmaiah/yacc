package auth

import (
	"encoding/json"
	"net/http"
	"yacc/backend/internal/db"
	"yacc/backend/internal/helpers"
)

type LoginRequest struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Email == "" || req.Passwd == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	session_id, err := db.Login(req.Email, req.Passwd)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{Name: "session_id", Value: session_id, MaxAge: 60 * 60 * 24 * 28, HttpOnly: true, Path: "/api", SameSite: http.SameSiteStrictMode}
	http.SetCookie(w, cookie)

}

func Signup(w http.ResponseWriter, r *http.Request) {

	var req LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	if req.Email == "" || req.Passwd == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	hash, err := helpers.Hash(req.Passwd)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	session_id, pgerr := db.Signup(req.Email, hash, helpers.RandUsername())

	if pgerr != nil {
		if pgerr.Code == "23505" {
			http.Error(w, "user already exists", http.StatusConflict)
		} else {
			http.Error(w, "internal error", http.StatusInternalServerError)
		}
		return
	}

	cookie := &http.Cookie{Name: "session_id", Value: session_id, MaxAge: 60 * 60 * 24 * 28, HttpOnly: true, Path: "/api", SameSite: http.SameSiteStrictMode}
	http.SetCookie(w, cookie)
}

func Logout(w http.ResponseWriter, r *http.Request) {

	session_cookie, err := r.Cookie("session_id")

	if err != nil || session_cookie.Value == "" {
		http.Error(w, "user not logged in", http.StatusUnauthorized)
		return
	}

	err = db.Logout(session_cookie.Value)
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}
	cookie := &http.Cookie{Name: "session_id", Value: "", MaxAge: -1, HttpOnly: true, Path: "/api", SameSite: http.SameSiteStrictMode}
	http.SetCookie(w, cookie)
}

func Verify(w http.ResponseWriter, r *http.Request) {
	session, err := r.Cookie("session_id")
	if err != nil {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	_, err = db.UserInfo(session.Value)
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
}

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
		return
	}

	db.ChangeUsername(username.Username, user_id)
}
