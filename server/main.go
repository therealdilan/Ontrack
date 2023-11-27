package main

import (
	"github.com/labstack/echo/v4"
	"github.com/therealdilan/ontrack/app"
)

func main() {
	e := echo.New()

	executeTemplate()

	e.Logger.Fatal(e.Start(":8000"))
}
