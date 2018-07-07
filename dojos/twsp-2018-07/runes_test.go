package main

import (
	"testing"
	"strings"
)

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}

const A = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`

func TestParseLine(t *testing.T) {
	char, name := ParseLine(A)
	wantChar := 'A'
	if char != wantChar {
		t.Errorf("Got: %q, want: %q", char, wantChar)
	}
	wantName := "LATIN CAPITAL LETTER A"
	if name != wantName {
		t.Errorf("Got: %q, want: %q", name, wantName)
	}
}

func TestFormatLine(t *testing.T) {
	got := FormatLine(ParseLine(A))
	want := "U+0041\tA\tLATIN CAPITAL LETTER A"
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}

const Cruzeiro = "20A2;CRUZEIRO SIGN;Sc;0;ET;;;;;N;;;;;"

func TestFormatLine_HexadecimalWithLetter(t *testing.T) {
	got := FormatLine(ParseLine(Cruzeiro))
	want := "U+20A2\t₢\tCRUZEIRO SIGN"
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}

}

const Data = `20A0;EURO-CURRENCY SIGN;Sc;0;ET;;;;;N;;;;;
20A1;COLON SIGN;Sc;0;ET;;;;;N;;;;;
20A2;CRUZEIRO SIGN;Sc;0;ET;;;;;N;;;;;`

func TestSearch(t *testing.T) {
	query := "CRUZEIRO"
	data := strings.NewReader(Data)
	got := Search(data,query)
	want := "U+20A2\t₢\tCRUZEIRO SIGN\n"
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}

func TestSearch_WithThreeResults(t *testing.T) {
	query := "SIGN"
	data := strings.NewReader(Data)
	got := Search(data,query)
	want := "U+20A0\t₠\tEURO-CURRENCY SIGN\nU+20A1\t₡\tCOLON SIGN\nU+20A2\t₢\tCRUZEIRO SIGN\n"
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}




















