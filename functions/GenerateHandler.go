package functions

import (
	"net/http"
	"strconv"
)

func GenerateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		RenderError(w, http.StatusMethodNotAllowed)
		return
	}

	banner := r.FormValue("banner")
	userText := r.FormValue("text")
	shouldDownload := r.FormValue("download")

	if banner == "" || userText == "" {
		RenderError(w, http.StatusBadRequest)
		return
	}

	if len(userText) > 3000 {
		RenderError(w, http.StatusBadRequest)
		return
	}

	asciiArt, status := GeneratingTheAsciiArt(w, banner, userText)

	if status != 200 {
		RenderError(w, status)
		return
	}

	type result struct {
		Text   string
		Result string
		Banner string
	}

	if shouldDownload != "true" {
		Exec("result.html", w, result{userText, asciiArt, banner})
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", "attachment")
	w.Header().Set("Content-Length", strconv.Itoa(len(asciiArt)))

	w.Write([]byte(asciiArt))
}
