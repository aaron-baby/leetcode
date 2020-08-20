package dynamic_programming

import "testing"

func TestCanCross(t *testing.T) {
	stones := []int{0, 1, 3, 5, 6, 8, 12, 17}
	got := canCross(stones)
	want := true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	stones = []int{0, 1, 2, 3, 4, 8, 9, 11}
	got = canCross(stones)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
