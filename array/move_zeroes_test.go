package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	want := []int{1, 3, 12, 0, 0}
	if !cmp.Equal(nums, want) {
		t.Errorf("got %v want %v given", nums, want)
	}
}
