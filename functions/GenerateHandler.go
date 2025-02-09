package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

// StatusInternalServerError           = 500
// StatusOK                   = 200
// StatusMethodNotAllowed             = 405
// StatusBadRequest                   = 400
func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderPageNotFound(w, http.StatusMethodNotAllowed)
		return
	}
	banner := r.FormValue("banner")
	userText := r.FormValue("text")
	if banner == "" || userText == "" {
		RenderPageNotFound(w, http.StatusBadRequest)
		return
	}
	if len(userText) > 3000 {
		RenderPageNotFound(w, http.StatusBadRequest)
		return
	}
	asciiArt, check := GeneratingTheAsciiArt(w, banner, userText)
	if check {
		return
	} else {
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
}
