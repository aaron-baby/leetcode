package dynamic_programming

import "testing"

func TestMinPathSum(t *testing.T) {
	grid := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}
	got := minPathSum(grid)
	want := 7
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
