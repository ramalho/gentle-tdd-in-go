package main

import (
	"testing"
	"strings"
)

const capitalA = `0041;LATIN CAPITAL LETTER A;Lu;0;L;;;;;N;;;;0061;`

func TestParseEntry(t *testing.T) {
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

func TestFilter(t *testing.T) {
	query := "LETTER A"
	dbFile := strings.NewReader(capitalA)
	results := Filter(dbFile, query)
	if len(results) != 1 {
		t.Errorf("Got %d results, expected 1", len(results))
	}
}

func TestMatch(t *testing.T) {
	query := "LETTER A"
	name := "LATIN CAPITAL LETTER A"
	want := true
	got := match(query, name)
	if got != want {
		t.Errorf("Match(%q, %q) got: %v, want: %v", query, name, got, want)
	}
}

func TestMatch_table(t *testing.T) {
	var testCases = []struct {
		query  	string
		name   	string
		want 	bool
	}{
		{"A", "A", true},
		{"A", "A B", true},
		{"A", "AB", false},
		{"A B", "A B", true},
		{"A B", "A B C", true},
		{"C A", "A B C", true},
		{"c A", "A B C", true},
		{"A B", "A B-C", true},
		{"A-B", "A B C", true},
	}
	for _, tc := range testCases {
		testName := tc.query + "⊆" + tc.name
		t.Run(testName, func(t *testing.T) {
			got := match(tc.query, tc.name)
			if tc.want != got {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}