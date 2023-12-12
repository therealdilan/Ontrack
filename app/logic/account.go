package logic

import (
	"fmt"
	"net/http"
  "firebase.google.com/go/v4/auth"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
    if Backend == nil {
      fmt.Println("Firebase backend is not initialized")
    return
  }

  client, err := Backend.Auth(ctx)

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



