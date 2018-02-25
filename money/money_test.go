package main

import (
	"testing"
)

func TestAmmount(t *testing.T) {
	five := Money{5, "USD"}
	x := 2
	got := five.Times(x)
	want := 10
	if got.Ammount != want {
		t.Errorf("five.times(%d), got: %d, want: %d.", x, got, want)
	}
}
