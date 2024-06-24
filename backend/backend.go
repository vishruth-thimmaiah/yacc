package backend

import (
	"cmp"
	"log"
	"net/http"
	"os"
)

func Init() {
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
