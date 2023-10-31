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
	ID      int
}

func main() {
	var todoList map[string]int

	fmt.Println("Server Working")

	todoList := make(map[string]task{
		"Tasks": {
			{Content: "Some content", State: false, ID: 1},
		},
	})

	addTask := func(w http.ResponseWriter, r *http.Request) {
		newTask := r.PostFormValue("newTask")
		htmlTemplate := template.Must(template.ParseFiles("index.html"))
		htmlTemplate.ExecuteTemplate(w, "taskItem", task{Content: newTask, State: false, ID: setID()})

	}

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setID() int {
	return len(todoList) + 1
}

//go run ./main.go
