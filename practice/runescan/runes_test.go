package main

import "testing"

func Example() {
	main()
	// Output:
	// Please provide one or more words to search.
}

func TestParseLine(t *testing.T) {
	line := "005A;LATIN CAPITAL LETTER Z;Lu;0;L;;;;;N;;;;007A;"
	got := ParseLine(line)
	want := CharName{'Z', "LATIN CAPITAL LETTER Z"}
	if got != want {
		t.Errorf("Got: %v, want: %v", got, want)
	}
}
