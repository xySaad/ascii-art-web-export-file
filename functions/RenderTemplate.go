package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(userText string, w http.ResponseWriter) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Could not load template", http.StatusInternalServerError)
		fmt.Println("Error parsing template:", err)
		return
	}
	err = tmpl.Execute(w, map[string]interface{}{
		"Text": userText,
	})
	if err != nil {
		http.Error(w, "Could not execute template", http.StatusInternalServerError)
		fmt.Println("Error executing template:", err)
		return
	}
}
