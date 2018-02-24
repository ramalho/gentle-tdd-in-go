package main

import (
	"testing"
)

func TestAmmount(t *testing.T) {
	five := Dollar{5}
	x := 2
	five.Times(x)
	want := 10
	if five.Ammount != want {
		t.Errorf("five.times(%d), got: %d, want: %d.", x, five.Ammount, want)
	}
}
