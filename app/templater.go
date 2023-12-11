package appHandler

import (
	"net/http"
	"text/template"
)

func HandleTemplate(route string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		t, err := template.ParseFiles("app/views/templates/"+route, "app/views/templates/navbar.html", "app/views/templates/login.html", "app/views/templates/content.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
    
    err = t.ExecuteTemplate(w, "app", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
