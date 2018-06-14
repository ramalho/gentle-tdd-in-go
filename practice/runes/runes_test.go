package main

import "testing"

func TestParseEntry(t *testing.T) {
	capitalA := `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`
	char, name := ParseEntry(capitalA)
	want := 'A'
	if char != want {
		t.Errorf("Wrong character: %q, expected: %q", char, want)
	}
	wantName := "LATIN CAPITAL LETTER A"
	if name != wantName {
		t.Errorf("Wrong name: %q, expected: %q", name, wantName)
	}
}

