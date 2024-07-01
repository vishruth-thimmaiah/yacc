package backend

import (
	"cmp"
	"go-msg/backend/internal/api"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Start(conn *pgxpool.Pool) {
	port := cmp.Or(os.Getenv("PORT"), "3000")

	// Frontend
	fs := http.FileServer(http.Dir("frontend/dist"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		fs.ServeHTTP(w, r)
	})
	http.Handle("/assets/", fs)

	// Start the server.
	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
