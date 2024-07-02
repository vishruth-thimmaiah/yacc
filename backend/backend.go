package backend

import (
	"cmp"
	"go-msg/backend/internal/api"
	"go-msg/backend/internal/db"
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

	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
