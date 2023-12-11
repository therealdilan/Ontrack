package main

import (
  "github.com/therealdilan/ontrack/app/logic"
	"github.com/therealdilan/ontrack/app"
)

func main() {
  appHandler.DefineRoutes() 
  logic.InitializeFirebaseApp()
}

