package logic

import (
	"fmt"
	"net/http"
  supa "github.com/nedpals/supabase-go"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
  user, err := App.Auth.SignUp(Ctx, supa.UserCredentials{
    Email:    "example@example.com",
    Password: "password",
  })

  if err != nil {
    panic(err)
  }

  fmt.Println(user)

  w.Write([]byte("works"))
}



