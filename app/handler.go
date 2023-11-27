package app

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func ReturnTemplate() {
	tmpl, _ = template.ParseFiles("views/templates/index.html")
}

func HandleTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl.Execute(w, nil)
}
