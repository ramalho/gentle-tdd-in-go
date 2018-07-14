package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/standupdev/strset"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide one or more words to search.")
		return
	}
	data, err := os.Open("UnicodeData.txt")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	query := strings.Join(os.Args[1:], " ")
	for _, cn := range Filter(data, query) {
		fmt.Printf("U+%04X\t%[1]c\t%s\n", cn.char, cn.name)
	}
}

// CharName struct holds a rune and its Unicode standard name.
type CharName struct {
	char rune
	name string
}

// parseLine extracts fields from a line in UnicodeData.txt.
func parseLine(line string) CharName {
	fields := strings.Split(line, ";")
	code, _ := strconv.ParseInt(fields[0], 16, 32)
	return CharName{rune(code), fields[1]}
}

// match tests whether a Set of query terms matches a character name.
func match(queryTerms strset.Set, name string) bool {
	return queryTerms.SubsetOf(strset.MakeFromText(name))
}

// Filter returns CharName records when query matches line in UnicodeData.txt
func Filter(data io.Reader, query string) []CharName {
	queryTerms := strset.MakeFromText(strings.ToUpper(query))
	scanner := bufio.NewScanner(data)
	result := []CharName{}
	for scanner.Scan() {
		charName := parseLine(scanner.Text())
		if match(queryTerms, charName.name) {
			result = append(result, charName)
		}
	}
	return result
}
