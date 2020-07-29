package dynamic_programming

import "testing"

func TestNumSquares(t *testing.T) {
	n := 12
	want := 3
	got := numSquares(n)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	got = numSquares(1)
	want = 1
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
