package app

import (
	"fmt"
	"os"
	"text/template"
)

type data struct {
	name string
}

func ExecuteTemplate() {
	p := data{name: "dilan"}

	t, error := template.New("views/templates/index.html").ParseFiles("views/templates")
	if error != nil {
		fmt.Print(error)
	}

	t.Execute(os.Stdout, p)
}
