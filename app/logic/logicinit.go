package logic

import (
  "context"
	"fmt"
  "os"

  "github.com/joho/godotenv"
  supa "github.com/nedpals/supabase-go"
	firebase "firebase.google.com/go/v4"
  "google.golang.org/api/option"
)

var (
  Ctx = context.Background()
  Fb *firebase.App
  Sb *supa.Client
)

func InitializeFirebaseApp() (*firebase.App, error) {
  opt := option.WithCredentialsFile("admin.firebase.json")
  app, err := firebase.NewApp(context.Background(), nil, opt)
  if err != nil {
    fmt.Println("Error initializing Firebase app:", err)
  }

  Fb = app

  fmt.Println(Fb)

  return Fb, nil
}

func InitializeSupabase() (*supa.Client, error) {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  SUPABASE_URL := os.Getenv("SUPABASE_URL")
  SUPABASE_KEY := os.Getenv("SUPABASE_KEY")

  supabase := supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

  Sb = supabase

  fmt.Println(Sb)

  return Sb, nil
}
