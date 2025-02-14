package functions

import (
	"net/http"
	"strings"
)

func GeneratingTheAsciiArt(w http.ResponseWriter, banner string, userText string) (string, bool) {
	for i := 0; i < len(userText); i++ {
		if !(userText[i] >= 32 && userText[i] <= 126) && !(strings.Contains(userText, "\r\n")) {
			RenderPageNotFound(w, http.StatusBadRequest)
			return "", true
		}
	}
	asciiArt := ""
	if banner == "Standard" {
		asciiArt = Mapping(banner, userText)
	} else if banner == "Shadow" {
		asciiArt = Mapping(banner, userText)
	} else if banner == "ThinkerToy" {
		asciiArt = Mapping(banner, userText)
	} else {
		RenderPageNotFound(w, http.StatusBadRequest)
		return "", true
	}
	return asciiArt, false
}
