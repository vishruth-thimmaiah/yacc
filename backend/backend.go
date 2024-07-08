package backend

import (
	"cmp"
	"log"
	"net/http"
	"os"
	"yacc/backend/internal/api"
	"yacc/backend/internal/api/auth"
	"yacc/backend/internal/db"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(conn *pgxpool.Pool) {

	port := cmp.Or(os.Getenv("PORT"), "3000")
	db.Setup(conn)
	auth.GOauth()
	// Frontend
	fs := http.FileServer(http.Dir("frontend/dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	http.Handle("/assets/", fs)

	// Backend
	http.HandleFunc("/api/auth/login", auth.Login)
	http.HandleFunc("/api/auth/google", auth.Google)
	http.HandleFunc("/api/auth/google/callback", auth.GoogleCb)
	http.HandleFunc("/api/auth/logout", auth.Logout)
	http.HandleFunc("/api/auth/signup", auth.Signup)

	http.HandleFunc("/api/user/contacts", api.Contacts)
	http.HandleFunc("/api/user/message", api.LoadMessages)
	http.HandleFunc("/api/send", api.SendMessage)
	http.HandleFunc("/api/addcontact", api.NewContact)

	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
