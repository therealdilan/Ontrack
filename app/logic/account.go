package logic

import (
  "context"
	"fmt"
	"net/http"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
  "google.golang.org/api/option"
)

var (
  ctx = context.Background()
  app *firebase.App
)

func InitializeFirebaseApp() (*firebase.App, error) {
  opt := option.WithCredentialsFile("admin.firebase.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Println("Error initializing Firebase app:", err)
  }
 
  return app, nil
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {
    if app == nil {
      fmt.Println("Firebase app is not initialized")
    return
  }

  client, err := app.Auth(ctx)

  if err != nil {
    fmt.Println("Error creating Firebase Auth client:", err)
    return
  }

  params := (&auth.UserToCreate{}).
          Email(r.FormValue("createAccountEmail")).
          EmailVerified(false).
          Password("password")

  user, err := client.CreateUser(ctx, params)

  if err != nil {
    fmt.Println("Error creating user:", err)
    return
  }

  fmt.Println("Successfully created user:", user)
}



