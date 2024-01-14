package app

import (
  "context"
	"fmt"
  "os"

  "github.com/joho/godotenv"
  supa "github.com/nedpals/supabase-go"
)

var (
  Ctx = context.Background()
  Supabase *supa.Client
)

func InitializeSupabase() (*supa.Client, error) {
  err := godotenv.Load()
  if err != nil {
    fmt.Println("Error loading .env file")
  }

  SUPABASE_URL := os.Getenv("SUPABASE_URL")
  SUPABASE_KEY := os.Getenv("SUPABASE_KEY")

  supabase := supa.CreateClient(SUPABASE_URL, SUPABASE_KEY)

  Supabase = supabase

  fmt.Println(Supabase)

  return Supabase, nil
}
