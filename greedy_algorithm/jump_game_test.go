package greedy_algorithm

import "testing"

func TestCanJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	got := canJump(nums)
	want := true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
