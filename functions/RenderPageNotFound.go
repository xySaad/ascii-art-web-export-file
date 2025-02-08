package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderPageNotFound(w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/page-not-found.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
	}
}
