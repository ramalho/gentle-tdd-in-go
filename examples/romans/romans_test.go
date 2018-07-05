package romans

import (
	"testing"
)

func TestRomanValue(t *testing.T) {
	var testCases = []struct {
		roman string
		want  int
	}{
		{"I", 1},
		{"II", 2},
		{"III", 3},
		{"IV", 4},
	}
	for _, tc := range testCases {
		t.Run(tc.roman, func(t *testing.T) {
			got := RomanValue(tc.roman)
			if got != tc.want {
				t.Errorf("got: %d, want: %d", got, tc.want)
			}
		})
	}
}
