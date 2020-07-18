package backtracking

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSolveNQueens(t *testing.T) {
	got := solveNQueens(2)
	want := [][]string{}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}

	// 4 dimensions
	got = solveNQueens(4)
	want = [][]string{
		{
			"..Q.",
			"Q...",
			"...Q",
			".Q..",
		},
		{
			".Q..",
			"...Q",
			"Q...",
			"..Q.",
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
