package strings_test

import (
    "strings"
    "testing"
    "fmt"
)

func TestIndex(t *testing.T) {
    const s, sep, want = "chicken", "ken", 4
    got := strings.Index(s, sep)
    if got != want {
        t.Errorf("Index(%q,%q) = %v; want %v", s, sep, got, want)
    }
}

func TestIndex_table(t *testing.T) {
    var testCases = []struct {
        s   string
        sub string
        want int
    }{
        {"", "", 0},
        {"", "a", -1},
        {"fo", "foo", -1},
        {"foo", "foo", 0},
        {"oofofoofooo", "f", 2},
        // etc
    }
    for _, tc := range testCases {
        got := strings.Index(tc.s, tc.sub)
        if got != tc.want {
            t.Errorf("Index(%q,%q) = %v; want %v", tc.s, tc.sub, got, tc.want)
        }
    }
}

func TestIndex_subtests(t *testing.T) {
    var testCases = []struct {
        s   string
        sub string
        want int
    }{
        {"", "", 0},
        {"", "a", -1},
        {"fo", "foo", -1},
        {"foo", "foo", 0},
        {"oofofoofooo", "f", 2},
        // etc
    }
    for _, tc := range testCases {
        t.Run(fmt.Sprintf("%q in %q", tc.sub, tc.s), func(t *testing.T) {
            got := strings.Index(tc.s, tc.sub)
            if got != tc.want {
                t.Errorf("Index(%q,%q) = %v; want %v", tc.s, tc.sub, got, tc.want)
            }
        })
    }
}




