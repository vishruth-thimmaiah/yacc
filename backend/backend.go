package backend

import (
	"cmp"
	"yacc/backend/internal/api"
	"yacc/backend/internal/db"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(conn *pgxpool.Pool) {

	port := cmp.Or(os.Getenv("PORT"), "3000")
	db.Setup(conn)
	// Frontend
	fs := http.FileServer(http.Dir("frontend/dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	http.Handle("/assets/", fs)

	// Backend
	http.HandleFunc("/api/user/login", api.Login)
	http.HandleFunc("/api/user/logout", api.Logout)
	http.HandleFunc("/api/user/signup", api.Signup)

	http.HandleFunc("/api/user/contacts", api.Contacts)
	http.HandleFunc("/api/user/message", api.LoadMessages)
	http.HandleFunc("/api/send", api.SendMessage)

	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
