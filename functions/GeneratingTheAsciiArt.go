package functions

import (
	"net/http"
)

func GeneratingTheAsciiArt(w http.ResponseWriter, banner string, userText string) (string, bool) {
	if !isClean(userText) {
		RenderErrPage(w, http.StatusBadRequest)
		return "", true
	}
	asciiArt := ""
	if banner == "Standard" {
		asciiArt = Mapping(banner, userText)
	} else if banner == "Shadow" {
		asciiArt = Mapping(banner, userText)
	} else if banner == "ThinkerToy" {
		asciiArt = Mapping(banner, userText)
	} else {
		RenderErrPage(w, http.StatusBadRequest)
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
