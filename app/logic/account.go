package logic

import (
	"fmt"
	"net/http"
  "firebase.google.com/go/v4/auth"
)

type CurrentUser struct {
	Email    string
	UID      string
	LoggedIn bool
}

func CheckUser(w http.ResponseWriter, r *http.Request) {
	client, err := Backend.Auth(ctx)

	if err != nil {
		fmt.Println("Error creating Firebase Auth client:", err)
		return
	}

	idToken := r.Header.Get("Authorization")
	if idToken == "" {
		fmt.Println("ID token not found in request headers")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := client.VerifyIDToken(ctx, idToken)

	if err != nil {
		fmt.Println("Error verifying ID token:", err)
		// Create an instance of CurrentUser and set LoggedIn to false
		currentUser := CurrentUser{LoggedIn: false}
		fmt.Println("Current user:", currentUser)
		return
	}

	// Create an instance of CurrentUser and set LoggedIn to true
	currentUser := CurrentUser{LoggedIn: true, Email: token.Claims["email"].(string), UID: token.UID}
	fmt.Println("Verified ID token:", token)
	fmt.Println("Current user:", currentUser)
}

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
    w.Write([]byte(err.Error()))
    return
  }

  fmt.Println("Successfully created user:", user)

  w.Write([]byte("Successfully created user"))
}

func LoginAccount(w http.ResponseWriter, r *http.Request) {
  if Backend == nil {
    fmt.Println("Firebase backend is not initialized")
    return
  }

  client, err := Backend.Auth(ctx)

  if err != nil {
    fmt.Println("Error creating Firebase Auth client:", err)
    return
  }

  user, err := client.GetUserByEmail(ctx, r.FormValue("loginEmail"))

  if err != nil {
    fmt.Println("Error getting user by email:", err)
    w.Write([]byte(err.Error()))
    return
  }

  fmt.Println("Successfully fetched user data:", user)

  w.Write([]byte("Successfully fetched user data"))
}

