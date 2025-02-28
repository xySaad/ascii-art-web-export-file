package functions

import (
	"html/template"
	"net/http"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderErrPage(w, http.StatusMethodNotAllowed)
		return
	}

	banner := r.FormValue("banner")
	userText := r.FormValue("text")
	if banner == "" || userText == "" {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}
	if len(userText) > 3000 {
		RenderErrPage(w, http.StatusBadRequest)
		return
	}
	asciiArt, check := GeneratingTheAsciiArt(w, banner, userText)
	if check {
		return
	} else {
		w.WriteHeader(http.StatusOK) // go sends 200 by default
		tmpl, err := template.ParseFiles("templates/result.html")
		if err != nil {
			RenderErrPage(w, http.StatusInternalServerError)
			return
		}
		err = tmpl.Execute(w, asciiArt)
		if err != nil {
			RenderErrPage(w, http.StatusInternalServerError)
			return
		}
	}
}
