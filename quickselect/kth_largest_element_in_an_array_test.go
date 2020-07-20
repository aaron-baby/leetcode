package quickselect

import "testing"

func TestFindKthLargest(t *testing.T) {
	nums, k := []int{3, 2, 1, 5, 6, 4}, 2
	want := 5
	got := findKthLargest(nums, k)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	nums, k = []int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4
	want = 4
	got = findKthLargest(nums, k)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
