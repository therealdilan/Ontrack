package main

import (
	"fmt"
	"net/http"

	"github.com/therealdilan/ontrack/app"
)

func main() {
	http.HandleFunc("/", app.HandleTemplate("index.html"))
	http.HandleFunc("/consolelog", consoleLog)
	http.ListenAndServe(":8080", nil)
}

func consoleLog(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GOTTEM"))
	fmt.Println("gottem")
}
