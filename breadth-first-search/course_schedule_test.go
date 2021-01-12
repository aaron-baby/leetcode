package breadth_first_search

import "testing"

func TestCanFinish(t *testing.T) {
	numCourses := 2
	prerequisites := [][]int{
		{1, 0},
		{0, 1},
	}
	got := canFinish(numCourses, prerequisites)
	want := false
	if got != want {
		t.Errorf("got %t want %t", got, want)
	}
}
