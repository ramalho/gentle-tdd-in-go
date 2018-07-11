package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Please provide one or more words to search.")
}

// CharName struct holds a rune and its Unicode standard name.
type CharName struct {
	char rune
	name string
}

// ParseLine extracts fields from a line in UnicodeData.txt.
func ParseLine(line string) CharName {
	fields := strings.Split(line, ";")
	code, _ := strconv.ParseInt(fields[0], 16, 32)
	return CharName{rune(code), fields[1]}
}
