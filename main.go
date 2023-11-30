package main

import (
	"fmt" 
	"net/http"

	"github.com/therealdilan/ontrack/app"
)

func main() {
	// Register the static route outside HandleTemplate
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))

	http.HandleFunc("/", app.HandleTemplate("index.html"))
	http.HandleFunc("/consolelog", consoleLog)
  
  http.ListenAndServe(":8080",nil)
}

func consoleLog(w http.ResponseWriter, r *http.Request) {
  fmt.Println("working :3")
  w.Write([]byte("working :3")) 
}
