package main

import (
	"strings"
	"strconv"
)

func ParseEntry(line string) (rune, string) {
	fields := strings.Split(line, ";")
	code, _ := strconv.ParseInt(fields[0], 16, 32)
	return rune(code), fields[1]
}
