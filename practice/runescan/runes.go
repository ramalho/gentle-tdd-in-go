package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Please provide one or more words to search.")
}

// ParseLine extracts fields from a line in UnicodeData.txt
func ParseLine(line string) string {
	fields := strings.Split(line, ";")
	return fields[1]
}
