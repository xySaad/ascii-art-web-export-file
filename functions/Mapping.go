package functions

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func Mapping(banner string, userText string) (result string, status int) {
	var content *os.File
	var err error

	content, err = os.Open("./banners/" + banner + ".txt")
	if err != nil {
		fmt.Printf("Error : %v\n", err)
		status = http.StatusInternalServerError
		return
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

	return result, http.StatusOK
}
