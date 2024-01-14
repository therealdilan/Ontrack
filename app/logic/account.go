package app

import (
	"fmt"
	"net/http"
  supa "github.com/nedpals/supabase-go"
)

func CreateUserAccount(w http.ResponseWriter, r *http.Request) {
  data := r.FormValue
  user, err := Supabase.Auth.SignUp(Ctx, supa.UserCredentials{
    Email:    data("userAccountEmail"),
    Password: data("userAccountPassword"),
  })

  if err != nil {
    panic(err)
  }

  fmt.Println(user)

  w.Write([]byte("works"))
}



