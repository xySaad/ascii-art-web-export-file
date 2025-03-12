package functions

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Mapping(banner string, userText string) string {
	var content *os.File
	var err error
	result := ""
	if banner == "Standard" {
		content, err = os.Open("standard.txt")
	} else if banner == "Shadow" {
		content, err = os.Open("shadow.txt")
	} else if banner == "ThinkerToy" {
		content, err = os.Open("thinkertoy.txt")
	}
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		os.Exit(1)
	}
	defer content.Close()
	ascii_map := make(map[int32]int32)
	var i int32 = 32

	for ; i <= 126; i++ {
		ascii_map[i] = ((i - 32) * 9) + 1
	}

	var Lines []string
	scanner := bufio.NewScanner(content)
	for scanner.Scan() {
		Lines = append(Lines, scanner.Text())
	}
	intoSlice := strings.Split(userText, "\n")
	for _, val := range intoSlice {
		result += Printing(val, Lines, ascii_map)
	}

	return result
}
