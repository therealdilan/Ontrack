package app

import (
	"fmt"
	"os"
	"text/template"
)

type data struct {
	name string
}

func executeTemplate() {
	p := data{name: "dilan"}

	t, error := template.New("templates/index.html").ParseFiles("templates")
	if error != nil {
		fmt.Print(error)
	}

	t.Execute(os.Stdout, p)
}
