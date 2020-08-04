package sliding_window

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	got := maxSlidingWindow(nums, 3)
	want := []int{3, 3, 5, 5, 6, 7}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
