package functions

import (
	"net/http"
	"slices"
)

func GeneratingTheAsciiArt(w http.ResponseWriter, banner string, userText string) (result string, success bool) {
	if !isClean(userText) {
		return
	}

	var banners = []string{"Standard", "Shadow", "ThinkerToy"}

	if !slices.Contains(banners, banner) {
		return
	}

	return Mapping(banner, userText), true
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
