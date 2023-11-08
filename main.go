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

var taskID int = 0
var tasks []task

func main() {

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)

	log.Fatal(http.ListenAndServe(":8000", nil))
	print("Server is running on port 8000")
}

func pageLoad(w http.ResponseWriter, r *http.Request) {
	htmlTemplate := template.Must(template.ParseFiles("index.html"))
	htmlTemplate.Execute(w, nil)
}

func addTask(w http.ResponseWriter, r *http.Request) {
	newTask := task{
		Content: r.FormValue("newTask"),
		State:   false,
		ID:      taskID,
	}

	tasks = append(tasks, newTask)
	taskID++

	htmlTemplate := template.Must(template.ParseFiles("index.html"))
	htmlTemplate.ExecuteTemplate(w, "taskItem", newTask)
	fmt.Println(tasks)
}
