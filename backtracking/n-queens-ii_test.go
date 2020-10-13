package backtracking

import "testing"

func TestTotalNQueens(t *testing.T) {
	n := 4
	got := totalNQueens(n)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
