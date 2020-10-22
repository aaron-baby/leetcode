package dynamic_programming

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}
	got := uniquePathsWithObstacles(obstacleGrid)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
