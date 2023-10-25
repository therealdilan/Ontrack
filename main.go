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
		newTask := r.PostFormValue("newTask")
		newTaskHTML := fmt.Sprintf("<div class='m-4 mt-4'><p class='text-2xl'>%s</p><button class='bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded'>Done</button><button class='bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded'>Remove</button></div>", newTask)
		htmlTemplate, _ := template.New("addNewTask").Parse(newTaskHTML)
		htmlTemplate.Execute(w, nil)
	}

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

//go run ./main.go
