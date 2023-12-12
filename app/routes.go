package appHandler

import (
  "net/http"
  "github.com/therealdilan/ontrack/app/logic"
)

func DefineApp() {
  logic.InitializeFirebaseApp()
  logic.InitializeSupabase()
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("app/static"))))
  http.HandleFunc("/", HandleTemplate("index.html"))
  // Start HTTP server
  http.ListenAndServe(":8080", nil)
}
