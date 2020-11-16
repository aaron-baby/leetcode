package main

import "testing"

func TestMyPow(t *testing.T) {
	x, n := 2.00000, 10
	got := myPow(x, n)
	want := 1024.00000
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	x, n = 2.00000, -2
	got = myPow(x, n)
	want = 0.25000
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
