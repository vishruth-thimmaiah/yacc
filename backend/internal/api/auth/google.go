package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"

	"golang.org/x/oauth2"
)

var googleOauth *oauth2.Config
var stateO string

func GOauth() {
	googleOauth = &oauth2.Config{
		RedirectURL:  "localhost:4000/api/auth/google/callback",
		ClientID:     os.Getenv("OAUTH_GOOGLE_CLIENT"),
		ClientSecret: os.Getenv("OAUTH_GOOGLE_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: oauth2.Endpoint{
			AuthURL:       "https://accounts.google.com/o/oauth2/auth",
			TokenURL:      "https://oauth2.googleapis.com/token",
			DeviceAuthURL: "https://oauth2.googleapis.com/device/code",
			AuthStyle:     oauth2.AuthStyleInParams,
		},
	}
}

func Google(w http.ResponseWriter, r *http.Request) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	stateO := base64.URLEncoding.EncodeToString(b)

	url := googleOauth.AuthCodeURL(stateO)

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCb(w http.ResponseWriter, r *http.Request) {
}
