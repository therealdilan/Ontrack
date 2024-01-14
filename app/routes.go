package appHandler

import (
  "net/http"
  "github.com/therealdilan/ontrack/app/logic"
)

func DefineApp() {
  app.InitializeSupabase()
  http.HandleFunc("/CreateUserAccount", app.CreateUserAccount) 
  //http.HandleFunc("/LoginUserAccount", logic.LoginUserAccount)
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))
  http.HandleFunc("/", HandleTemplate("index.html"))
  // Start HTTP server
  http.ListenAndServe(":8080", nil)
}
