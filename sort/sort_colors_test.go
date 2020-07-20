package sort

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSortColors(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	want := []int{0, 0, 1, 1, 2, 2}
	if !cmp.Equal(nums, want) {
		t.Errorf("got %v want %v given", nums, want)
	}

	nums1 := []int{2, 0, 1}
	sortColors(nums1)
	want = []int{0, 1, 2}
	if !cmp.Equal(nums1, want) {
		t.Errorf("got %v want %v given", nums1, want)
	}
}
