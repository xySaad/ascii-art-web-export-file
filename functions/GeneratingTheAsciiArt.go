package functions

import (
	"errors"
	"strings"
)

func GeneratingTheAsciiArt(banner string, userText string) (string, int, error) {
	for i := 0; i < len(userText); i++ {
		if !(userText[i] >= 32 && userText[i] <= 126) && !(strings.Contains(userText, "\r\n")) {
			return "", 1, errors.New("⚠️ Warning: Input contains non printable ASCII characters")
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
		return "", 2, errors.New("⚠️ Warning: Please Select A Banner")
	}
	return asciiArt, 0, nil
}
