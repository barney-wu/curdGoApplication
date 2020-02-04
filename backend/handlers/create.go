package handlers

import (
	"fmt"
	"net/http"
	"text/template"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("CreateHandler path", r.URL.Path)
	var tmpl = template.Must(template.ParseGlob("frontend/*"))
	tmpl.ExecuteTemplate(w, "Create", nil)
}
