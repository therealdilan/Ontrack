package main

import (
  "fmt"
  "context"
  "net/http"

  firebase "firebase.google.com/go"

  "google.golang.org/api/option"
  "github.com/therealdilan/ontrack/app"
)

func main() {
  opt := option.WithCredentialsFile("admin.firebase.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  
  fmt.Println("backend: ", app)

  if err != nil {
    fmt.Errorf("error initializing app: ", err)
  } // Register the static route outside HandleTemplate
	
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))
	http.HandleFunc("/", handler.HandleTemplate("index.html"))
  http.ListenAndServe(":8080",nil)
}

func consoleLog(w http.ResponseWriter, r *http.Request) {
  fmt.Println("working :3")
  w.Write([]byte("working :3")) 
}
