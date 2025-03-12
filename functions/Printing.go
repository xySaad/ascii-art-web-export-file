package functions

func Printing(str string, Lines []string, ascii_map map[int32]int32) string {
	charRows := [][]string{}
	result := ""
	if len(str) == 0 {
		result = "\n"
		return result
	}
	for _, char := range str {
		startLine := ascii_map[char]
		charArt := Lines[startLine : startLine+9]
		charRows = append(charRows, charArt)
	}
	s := ""
	for i := 0; i < 8; i++ {
		s = ""
		for _, val := range charRows {
			s += val[i]
		}
		result += s + "\n"
	}
	return result
}
