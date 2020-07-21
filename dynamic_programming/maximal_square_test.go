package dynamic_programming

import "testing"

func TestMaximalSquare(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}
	got := maximalSquare(matrix)
	want := 4
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
