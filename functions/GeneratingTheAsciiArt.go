package functions

import (
	"net/http"
	"slices"
)

func GeneratingTheAsciiArt(w http.ResponseWriter, banner string, userText string) (result string, status int) {
	if !isClean(userText) {
		status = http.StatusBadRequest
		return
	}

	banners, status := ListBanners()
	if status != 200 {
		return
	}

	if !slices.Contains(banners, banner) {
		status = http.StatusBadRequest
		return
	}
	return Mapping(banner, userText)
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
