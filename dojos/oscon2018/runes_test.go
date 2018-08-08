package main

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
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

func TestParse(t *testing.T) {
	// Given
	const line = "0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;"

	// When
	gotCode, gotName := Parse(line)

	// Then
	wantCode := "0041"
	if wantCode != gotCode {
		t.Errorf("code: %v, want: %v", gotCode, wantCode)
	}
	wantName := "LATIN CAPITAL LETTER A"
	if wantName != gotName {
		t.Errorf("name: %v, want: %v", gotName, wantName)
	}
}

func TestMatch(t *testing.T) {
	testCases := []struct {
		query string
		name  string
		want  bool
	}{
		{"BLACK", "BLACK CHESS KING", true},
		{"WHITE", "BLACK CHESS KING", false},
		{"black", "BLACK CHESS KING", true},
		{"chess black", "BLACK CHESS KING", true},
		{"", "BLACK CHESS KING", false},
		{"BLACK CAT", "BLACK CHESS KING", false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v in %v", tc.query, tc.name),
			func(t *testing.T) {
				// Given
				query := tc.query
				name := tc.name

				// When
				got := Match(query, name)
				if got != tc.want {
					t.Errorf("got %v; want %v", got, tc.want)
				}
			})
	}
}

const sample = `003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
`

func TestSelect(t *testing.T) {
	// Given
	query := "LETTER"
	data := strings.NewReader(sample)
	// When
	got := Select(data, query)
	// Then
	want := []CodeName{
		{"0041", "LATIN CAPITAL LETTER A"},
		{"0042", "LATIN CAPITAL LETTER B"},
	}
	assert.Equal(t, want, got)
}

func TestStringToRune(t *testing.T) {
	// GIVEN
	code := "0041"
	// When
	got := StringToRune(code)
	// Then
	assert.Equal(t, 'A', got)
}
