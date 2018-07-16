package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/standupdev/strset"
	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}

func TestParseLine(t *testing.T) {
	// given
	line := "005A;LATIN CAPITAL LETTER Z;Lu;0;L;;;;;N;;;;007A;"
	// when
	got := parseLine(line)
	// then
	want := CharName{'Z', "LATIN CAPITAL LETTER Z"}
	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
func TestMatch(t *testing.T) {
	testCases := []struct {
		query string
		name  string
		want  bool
	}{
		{"BLACK", "BLACK CHESS KING", true},
		{"CHESS BLACK", "WHITE CHESS KING", false},
		{"CHESS BLACK", "BLACK CHESS KING", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v in %v", tc.query, tc.name), func(t *testing.T) {
			// given
			queryTerms := strset.MakeFromText(tc.query)
			// when
			got := match(queryTerms, tc.name)
			// then
			if got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

const dataSample = `002D;HYPHEN-MINUS;Pd;0;ES;;;;;N;;;;;
002E;FULL STOP;Po;0;CS;;;;;N;PERIOD;;;;
002F;SOLIDUS;Po;0;CS;;;;;N;SLASH;;;;
0030;DIGIT ZERO;Nd;0;EN;;0;0;0;N;;;;;
0031;DIGIT ONE;Nd;0;EN;;1;1;1;N;;;;;
0032;DIGIT TWO;Nd;0;EN;;2;2;2;N;;;;;
`

func TestFilter(t *testing.T) {
	// given
	data := strings.NewReader(dataSample)
	query := "DIGIT"
	results := make(chan CharName)
	// when
	go Filter(data, query, results)
	// then
	expected := []CharName{
		{'0', "DIGIT ZERO"},
		{'1', "DIGIT ONE"},
		{'2', "DIGIT TWO"},
	}
	for _, want := range expected {
		got := <-results
		assert.Equal(t, want, got)
	}
	res, read := <-results
	assert.False(t, read, fmt.Sprintf("unexpected result: %#v", res))
}

func Example_chess_black() {
	savedArgs := os.Args
	defer func() { os.Args = savedArgs }()
	os.Args = []string{"", "chess black"}
	main()
	// Output:
	// U+265A	♚	BLACK CHESS KING
	// U+265B	♛	BLACK CHESS QUEEN
	// U+265C	♜	BLACK CHESS ROOK
	// U+265D	♝	BLACK CHESS BISHOP
	// U+265E	♞	BLACK CHESS KNIGHT
	// U+265F	♟	BLACK CHESS PAWN
}
