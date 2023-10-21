package main 

import (
	"fmt"
	"log"
	"html/template"
	"net/http"
)

type task struct {
	Title string
	Content string
}

func main() {
	fmt.Println("Hello World")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		htmlTemplate := template.Must(template.ParseFiles("index.html"))
		htmlTemplate.Execute(w, nil)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

