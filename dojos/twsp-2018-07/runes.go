package main

import (
	"fmt"
	"strings"
	"strconv"
	"bufio"
	"os"
	"io"
)

func ParseLine(line string) (rune, string) {
	fields := strings.Split(line, ";")
	code, _ := strconv.ParseInt(fields[0], 16, 32)
	return rune(code), fields[1]
}

func FormatLine(char rune, name string) string {
	return fmt.Sprintf("U+%04X\t%c\t%v", char, char, name)
}

func Search(data io.Reader, query string) string {
	scanner := bufio.NewScanner(data)
	builder := strings.Builder{}
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue

		}
		if strings.Contains(line, strings.ToUpper(query)) {
			builder.WriteString(FormatLine(ParseLine(line)))
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide one or more words to search.")
		return
	}

	data, err := os.Open("UnicodeData.txt")

	if err != nil {
		panic(err)
	}

	result := Search(data, strings.Join(os.Args[1:], " "))

	fmt.Print(result)




}











