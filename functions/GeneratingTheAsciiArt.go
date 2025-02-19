package functions

import (
	"net/http"
	"strings"
)

func GeneratingTheAsciiArt(w http.ResponseWriter, banner string, userText string) (string, bool) {
	if !isClean(userText) {
		RenderPageNotFound(w, http.StatusBadRequest)
		return "", true
	}
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

func isClean(str string) bool {
	for i := 0; i < len(str); i++ {
		if (str[i] == '\r' && str[i+1] == '\n') || str[i] == '\n' {
			continue
		}
		if str[i] < 32 || str[i] > 126 {
			return false
		}
	}
	return true
}
