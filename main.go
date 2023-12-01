package main

import (
	"net/http"

	"github.com/therealdilan/ontrack/app"
)

func main() {
	// Register the static route outside HandleTemplate
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))

	http.HandleFunc("/", app.HandleTemplate("index.html"))  
  http.ListenAndServe(":8080",nil)
}

