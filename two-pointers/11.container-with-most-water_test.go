package two_pointers

import "testing"

func TestMaxArea(t *testing.T) {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	got := maxArea(height)
	want := 49
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
