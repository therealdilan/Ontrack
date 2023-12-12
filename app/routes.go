package appHandler

import (
  "net/http"
  "github.com/therealdilan/ontrack/app/logic"
)

func DefineApp() {
  http.HandleFunc("/createAccount", logic.CreateAccount)
  http.HandleFunc("/loginAccount", logic.LoginAccount)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))
  http.HandleFunc("/", HandleTemplate("index.html"))
  // Start HTTP server
  http.ListenAndServe(":8080", nil)
}
