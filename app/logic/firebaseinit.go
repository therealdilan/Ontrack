package logic

import (
  "context"
	"fmt"

	firebase "firebase.google.com/go/v4"
  "google.golang.org/api/option"
)

var (
  ctx = context.Background()
  Backend *firebase.App
)

func InitializeFirebaseApp() (*firebase.App, error) {
  opt := option.WithCredentialsFile("admin.firebase.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Println("Error initializing Firebase app:", err)
  }

  Backend = app

  fmt.Println(Backend)

  return Backend, nil
}
