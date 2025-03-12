package functions

import (
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		RenderError(w, http.StatusNotFound)
		return
	}

	if Exec("index.html", w, nil) != nil {
		RenderError(w, http.StatusInternalServerError)
		return
	}
}
