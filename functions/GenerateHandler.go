package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	banner := r.FormValue("banner")
	userText := r.FormValue("text")
	if len(userText) > 3000 {
		http.Error(w, "Payload Too Large: text exceeds 1000 characters", http.StatusRequestEntityTooLarge)
		return
	}
	asciiArt, j, err := GeneratingTheAsciiArt(banner, userText)
	if err != nil {
		ErrHandling(err, j, w)
		RenderTemplate(userText, w)
		return
	}
	tmpl, err := template.ParseFiles("templates/result.html")
	if err != nil {
		fmt.Println("Error loading template:", err)
		return
	}
	err = tmpl.Execute(w, asciiArt)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
