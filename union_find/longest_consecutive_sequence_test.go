package union_find

import "testing"

func TestLongestConsecutive(t *testing.T) {
	nums := []int{100, 4, 200, 1, 3, 2}
	want := 4
	got := longestConsecutive(nums)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
