package dynamic_programming

import "testing"

func TestCanPartition(t *testing.T) {
	nums := []int{1, 5, 11, 5}
	got := canPartition(nums)
	want := true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	nums = []int{1, 2, 3, 5}
	got = canPartition(nums)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
