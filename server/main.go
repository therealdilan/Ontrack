package main

import (
	"net/http"

	"github.com/therealdilan/ontrack/app"
)

func main() {
	// Initialize the template
	app.ReturnTemplate()
	http.HandleFunc("/", app.HandleTemplate)
	http.ListenAndServe(":8000", nil)
}
