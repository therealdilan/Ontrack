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
  Backend, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Println("Error initializing Firebase app:", err)
  }
  
  return Backend, nil
}
