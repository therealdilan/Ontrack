package main

import (
	"net/http"

	"github.com/therealdilan/ontrack/app"
)

func main() {
	http.HandleFunc("/", app.HandleTemplate("lol.html"))
	http.ListenAndServe(":8080", nil)
}
