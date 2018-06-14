package main

import (
	"strings"
	"strconv"
	"io"
	"github.com/standupdev/strset"
)

func ParseEntry(line string) (rune, string) {
	fields := strings.Split(line, ";")
	code, _ := strconv.ParseInt(fields[0], 16, 32)
	return rune(code), fields[1]
}

func match(query, name string) bool {
	normalized := strings.Replace(strings.ToUpper(query), "-", " ", -1)
	queryTerms := strset.MakeFromText(normalized)
	nameTerms := strset.MakeFromText(strings.Replace(name, "-", " ", -1))
	return queryTerms.SubsetOf(nameTerms)
}

func Filter(dbFile io.Reader, query string) []rune {
	return []rune{'A'}
}
