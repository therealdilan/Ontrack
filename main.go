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

	pageLoad := func(w http.ResponseWriter, r *http.Request) {
		htmlTemplate := template.Must(template.ParseFiles("index.html"))

		todoList := map[string][]task{
			"Tasks": {
				{Content: "Some content", State: true},
				{Content: "Some more content", State: false},
			},
		}

		htmlTemplate.Execute(w, todoList)
	}

	addTask := func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.Header.Get("HX-Request"))
	}

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//go run ./main.go
