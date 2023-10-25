package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type task struct {
	State   bool
	Content string
}

func main() {
	fmt.Println("Server Working")

	h1 := func(w http.ResponseWriter, r *http.Request) {
		htmlTemplate := template.Must(template.ParseFiles("index.html"))
		htmlTemplate.Execute(w, nil)

		todoList := map[string][]task{
			"Tasks": {
				{Content: "Some content", State: true},
				{Content: "Some more content", State: false},
			},
		}

		htmlTemplate.Execute(w, todoList)
	}

	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
