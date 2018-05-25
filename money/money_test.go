package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestTimes(t *testing.T) {
	five := Money{5, "USD"}
	x := 2
	got := five.Times(x)
	want := Money{10, "USD"}
	if got != want {
		t.Errorf("five.times(%d), got: %v, want: %v.", x, got, want)
	}
}


func TestEquals(t *testing.T) {
	testCases := []struct {
		a Money
		b Money
		want  bool
	}{
		{Money{5, "USD"}, Money{5, "USD"}, true},
		{Money{5, "USD"}, Money{7, "USD"}, false},
		{Money{5, "USD"}, Money{5, "BRL"}, false},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v.Equal(%v) is %v", tc.a, tc.b, tc.want), func(t *testing.T) {
			got := tc.a.Equal(tc.b)
			assert.Equal(t, tc.want, got)
		})
	}
}
