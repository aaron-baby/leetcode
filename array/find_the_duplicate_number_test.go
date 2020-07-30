package array

import "testing"

func TestFindDuplicate(t *testing.T) {
	nums := []int{1, 3, 4, 2, 2}
	got := findDuplicate(nums)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
