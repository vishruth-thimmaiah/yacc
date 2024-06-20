package backend

import (
	"cmp"
	"fmt"
	"log"
	"net/http"
	"os"
)

func Init() {
	port := cmp.Or(os.Getenv("PORT"), "3000")

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("./frontend/dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":"+port, nil),
	)
}
