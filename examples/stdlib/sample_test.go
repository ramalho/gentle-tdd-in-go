package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestParseInt(t *testing.T) {
	input := "2A"
	want := int64(42)
	got, _ := strconv.ParseInt(input, 16, 7)
	if got != want {
		t.Errorf("Parsing %q, got: %d, want: %d", input, got, want)
	}
}

func TestParseInt_withErrorHandling(t *testing.T) {
	input := "2A"
	want := int64(42)
	got, err := strconv.ParseInt(input, 16, 7)
	if err != nil {
		t.Fatal(err)
	}
	if got != want {
		t.Errorf("Parsing %q, got: %d, want: %d", input, got, want)
	}
}

func EqualSlices(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, item := range s1 {
		if item != s2[i] {
			return false
		}
	}
	return true
}

func TestEqualSlices(t *testing.T) {
	var testCases = []struct {
		s1   []string
		s2   []string
		want bool
	}{
		{[]string{"A"}, []string{"A"}, true},
		{[]string{"A", "B", "C"}, []string{"A", "B", "C"}, true},
		{[]string{}, []string{}, true},
		{[]string{"A"}, []string{"B"}, false},
		{[]string{"A"}, []string{"A", "A"}, false},
		{[]string{"A", "B", "C"}, []string{"A", "B", "D"}, false},
		{[]string{"A"}, []string{}, false},
		{[]string{}, []string{"A"}, false},
	}
	for _, tc := range testCases {
		testName := fmt.Sprint(tc.s1, tc.s2)
		t.Run(testName, func(t *testing.T) {
			got := EqualSlices(tc.s1, tc.s2)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestSplitN(t *testing.T) {
	var testCases = []struct {
		text  string
		sep   string
		parts int
		want  []string
	}{
		{"A-B-C-D", "-", -1, []string{"A", "B", "C", "D"}},
		{"A-B-C-D", "|", -1, []string{"A-B-C-D"}},
		{"", "-", -1, []string{""}},
		{"A-B-C-D", "-", 0, []string{}},
		{"", "-", 0, []string{}},
		{"A-B-C-D", "-", 2, []string{"A", "B-C-D"}},
		{"A-B-C-D", "-", 3, []string{"A", "B", "C-D"}},
		{"A-B-C-D", "-", 8, []string{"A", "B", "C", "D"}},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("%q,%q,%d", tc.text, tc.sep, tc.parts)
		t.Run(testName, func(t *testing.T) {
			got := strings.SplitN(tc.text, tc.sep, tc.parts)
			if !EqualSlices(tc.want, got) {
				t.Errorf("got: %#v, want: %#v", got, tc.want)
			}
		})
	}
}
