package functions

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderErrPage(w, http.StatusNotFound)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, nil)
	if err != nil {
		RenderErrPage(w, http.StatusInternalServerError)
		return
	}
}
