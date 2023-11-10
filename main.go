package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type task struct {
	State   bool
	Content string
	ID      int
}

func (t task) String() string {
	return fmt.Sprintf("ID: %d, Content: %s, State: %t", t.ID, t.Content, t.State)
}

var taskID int = 0
var tasks []task

func main() {

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)
	http.HandleFunc("/remove-task", removeTask)
	http.HandleFunc("/mark-task", markTask)

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

func removeTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // Get the ID from the URL query parameter
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	removeTaskByID(id)
}

func removeTaskByID(id int) {
	for i, t := range tasks {
		if t.ID == id {
			// Remove the task from the slice by slicing it
			tasks = append(tasks[:i], tasks[i+1:]...)
			return
		}
	}
	fmt.Println(tasks)
}

func markTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // Get the ID from the URL query parameter
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	markTaskasDone(id, w)
	fmt.Println(tasks)
}

func markTaskasDone(id int, w http.ResponseWriter) {
	var taskToUpdate *task
	for i := range tasks {
		if tasks[i].ID == id {
			taskToUpdate = &tasks[i]
			break
		}
	}
	if taskToUpdate == nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	taskToUpdate.State = !taskToUpdate.State
	htmlTemplate := template.Must(template.ParseFiles("index.html"))
	htmlTemplate.ExecuteTemplate(w, "taskItem", taskToUpdate)
	fmt.Println(tasks)
}
