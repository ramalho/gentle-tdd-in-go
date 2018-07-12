package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/standupdev/strset"
	"github.com/stretchr/testify/assert"
)

func Example() {
	main()
	// Output:
	// Please enter one or more words to search.
}

func TestParseLine(t *testing.T) {
	// Given 0030;DIGIT ZERO;Nd;0;EN;;0;0;0;N;;;;;
	wantName := "DIGIT ZERO"
	wantChar := '0'
	input := "0030;DIGIT ZERO;Nd;0;EN;;0;0;0;N;;;;;"
	// When
	gotName, gotChar := parseLine(input)
	// Then
	if gotName != wantName {
		t.Errorf("Name => Got: %#v, want: %#v", gotName, wantName)
	}
	if gotChar != wantChar {
		t.Errorf("Char => Got: %#v, want: %#v", gotChar, wantChar)
	}
}

func TestMatch_Table(t *testing.T) {
	testCases := []struct {
		query strset.Set
		name  string
		want  bool
	}{
		{strset.MakeFromText("CHESS BLACK"), "WHITE CHESS KING", false},
		{strset.MakeFromText("CHESS BLACK"), "BLACK CHESS KING", true},
		{strset.MakeFromText("BLACK"), "BLACK CHESS KING", true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v in %v", tc.query, tc.name), func(t *testing.T) {
			if got := match(tc.query, tc.name); got != tc.want {
				t.Errorf("got %v; want %v", got, tc.want)
			}
		})
	}
}

const dataStr = `003C;LESS-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003D;EQUALS SIGN;Sm;0;ON;;;;;N;;;;;
003E;GREATER-THAN SIGN;Sm;0;ON;;;;;Y;;;;;
003F;QUESTION MARK;Po;0;ON;;;;;N;;;;;
0040;COMMERCIAL AT;Po;0;ON;;;;;N;;;;;
0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;
0042;LATIN CAPITAL LETTER B;Lu;0;L;;;;;N;;;;0062;
`

func TestFilter(t *testing.T) {
	query := "sign"
	data := strings.NewReader(dataStr)
	got := Filter(data, query)
	want := []string{
		"U+003C\t<\tLESS-THAN SIGN",
		"U+003D\t=\tEQUALS SIGN",
		"U+003E\t>\tGREATER-THAN SIGN",
	}
	assert.Equal(t, want, got)
}
