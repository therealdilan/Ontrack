package app

import (
	"os"
	"text/template"
)

func executeTemplate() {
	t, _ := template.ParseFiles(
		"templates/index.gohtml",
	)

	t.Execute(os.Stdout, map[string]string{
		"name": "dilan",
	})
}
