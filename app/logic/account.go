package logic 

import (
  "fmt"
  "net/http"

  supa "github.com/nedpals/supabase-go"
)

func CreateUserAccount(w http.ResponseWriter, r *http.Request) {
  user, err := Sb.Auth.SignUp(Ctx, supa.UserCredentials{
    Email:     r.FormValue("userAccountEmail"),
    Password: r.FormValue("userAccountPassword"),
  })

  if err != nil {
    w.Write([]byte(err.Error()))
    return
  }

  w.Write([]byte("User Created!"))

  fmt.Println(user) 
}

func LoginUserAccount(w http.ResponseWriter, r *http.Request) {
  user, err := Sb.Auth.SignIn(Ctx, supa.UserCredentials{
    Email:     r.FormValue("userAccountEmail"),
    Password: r.FormValue("userAccountPassword"),
  })

  if err != nil {
    w.Write([]byte(err.Error()))
    return
  }

  w.Write([]byte("User Logged In!"))

  fmt.Println(user)
}
