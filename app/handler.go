package app

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

var tmpl *template.Template

func ReturnTemplate() {
	tmpl, _ = template.ParseFiles("../app/views/templates/index.html")
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Current Working Directory:", currentDir)

}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
