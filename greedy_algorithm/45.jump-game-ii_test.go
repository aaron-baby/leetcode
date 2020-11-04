package greedy_algorithm

import "testing"

func TestJump(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	got := jump(nums)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
