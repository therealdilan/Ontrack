package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve the static files from the "public" directory
	fs := http.FileServer(http.Dir("../app/public/"))
	http.Handle("/", fs)

	http.HandleFunc("/change", changeText)

	// Start the server
	log.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func changeText(w http.ResponseWriter, r *http.Request) {
	newText := "Changed!!!"

	w.Write([]byte(newText))
}
