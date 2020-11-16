package hash_table

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	got := twoSum(nums, target)
	want := []int{0, 1}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
