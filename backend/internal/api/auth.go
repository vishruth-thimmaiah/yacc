package api

import (
	"go-msg/backend/internal/db"
	"go-msg/backend/internal/helpers"
	"net/http"
)

type LoginRequest struct {
	Email  string `json:"email"`
	Passwd string `json:"passwd"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	loginDetails := LoginRequest{Email: r.FormValue("email"), Passwd: r.FormValue("passwd")}

	if loginDetails.Email == "" || loginDetails.Passwd == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	session_id, err := db.Login(loginDetails.Email, loginDetails.Passwd)
	if err != nil {
		http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return
	}

	cookie := &http.Cookie{Name: "session_id", Value: session_id, HttpOnly: true}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/messages", http.StatusPermanentRedirect)

}

func Logout(w http.ResponseWriter, r *http.Request) {

	session_cookie, err := r.Cookie("session_id")

	if err != nil || session_cookie.Value == "" {
		http.Error(w, "user not logged in", http.StatusBadRequest)
		return
	}

	cookie := &http.Cookie{Name: "session_id", Value: "", MaxAge: -1}
	http.SetCookie(w, cookie)
}

func Signup(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}
	loginDetails := LoginRequest{Email: r.FormValue("email"), Passwd: r.FormValue("passwd")}

	if loginDetails.Email == "" || loginDetails.Passwd == "" {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	hash, err := helpers.Hash(loginDetails.Passwd)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	session_id, err := db.Signup(loginDetails.Email, hash)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	cookie := &http.Cookie{Name: "session_id", Value: session_id, HttpOnly: true}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/messages", http.StatusPermanentRedirect)

}
