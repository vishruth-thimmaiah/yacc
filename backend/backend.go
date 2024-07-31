package backend

import (
	"cmp"
	"log"
	"net/http"
	"os"
	"yacc/backend/internal/api"
	"yacc/backend/internal/api/auth"
	"yacc/backend/internal/api/messages"
	"yacc/backend/internal/db"
	"yacc/backend/internal/s3"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(conn *pgxpool.Pool) {

	port := cmp.Or(os.Getenv("PORT"), "3000")
	db.Setup(conn)
	bucket.Setup()
	auth.GOauth()

	// Frontend
	fs := http.FileServer(http.Dir("frontend/dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	http.Handle("/assets/", fs)

	// Backend

	hub := messages.CreateHub()
	go hub.Run()

	http.HandleFunc("/api/auth/login", auth.Login)
	http.HandleFunc("/api/auth/google", auth.Google)
	http.HandleFunc("/api/auth/google/callback", auth.GoogleCb)
	http.HandleFunc("/api/auth/signup", auth.Signup)
	http.HandleFunc("/api/auth/verify", auth.Verify)
	http.HandleFunc("/api/auth/logout", auth.Logout)
	http.HandleFunc("/api/auth/username", auth.ChangeUsername)
	http.HandleFunc("/api/auth/delete", auth.DeleteAccount)

	http.HandleFunc("/api/user/contacts", api.Contacts)
	http.HandleFunc("/api/user/message", messages.LoadMessages)
	http.HandleFunc("/api/addcontact", api.NewContact)

	http.HandleFunc("/api/messages", func(w http.ResponseWriter, r *http.Request) {
		messages.Messages(hub, w, r)
	})
	http.HandleFunc("/api/messages/upload", messages.Attachments)

	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
