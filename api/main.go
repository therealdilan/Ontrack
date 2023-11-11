package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type task struct {
	State   bool
	Content string
	ID      int
}

func (t task) String() string { // For Debugging
	return fmt.Sprintf("ID: %d, Content: %s, State: %t", t.ID, t.Content, t.State)
}

var taskID int = 0
var tasks []task

func main() {

	// Every HTTP request handling

	http.HandleFunc("/", pageLoad)
	http.HandleFunc("/add-task", addTask)
	http.HandleFunc("/remove-task", removeTask)
	http.HandleFunc("/mark-task", markTask)

	// Start the server

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello, Railway!",
		})
	})

	app.Listen(getPort())
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	return port
}

func pageLoad(w http.ResponseWriter, r *http.Request) { // Load the main page
	htmlTemplate := template.Must(template.ParseFiles("index.html"))
	htmlTemplate.Execute(w, nil)
}

func addTask(w http.ResponseWriter, r *http.Request) { // Adding a new Task
	newTask := task{
		Content: r.FormValue("newTask"),
		State:   false,
		ID:      taskID,
	}
	tasks = append(tasks, newTask) // Append the new task to the array
	taskID++                       // Increment the ID to be used for the new added task

	htmlTemplate := template.Must(template.ParseFiles("index.html"))
	htmlTemplate.ExecuteTemplate(w, "taskItem", newTask)
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // Get the ID from the URL query parameter
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam) // Convert the ID to an integer
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	removeTaskByID(id)
}

func removeTaskByID(id int) {
	for i, t := range tasks { // Find the task with the given ID
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...) // Removing the task
			return
		}
	}
	fmt.Println(tasks)
}

func markTask(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get("id") // Get the ID from the URL query parameter
	fmt.Println(idParam)
	id, err := strconv.Atoi(idParam) // Convert the ID to an integer
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	markTaskasDone(id, w) // Sending the new HTML with the changed task
}

func markTaskasDone(id int, w http.ResponseWriter) {
	var taskToUpdate *task
	for i := range tasks { // Find the task with the given ID
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

	htmlTemplate.ExecuteTemplate(w, "taskItem", taskToUpdate) // Once you get the task, send the task (with new State)
}
