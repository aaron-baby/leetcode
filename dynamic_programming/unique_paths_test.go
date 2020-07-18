package dynamic_programming

import "testing"

func TestUniquePaths(t *testing.T) {
	got := uniquePaths(3, 2)
	want := 3
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	got = uniquePaths(7, 3)
	want = 28
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	got = uniquePaths(51, 9)
	want = 1916797311
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
