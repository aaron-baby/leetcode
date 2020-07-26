package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5, 7, 7, 8, 8, 10}
	target := 8
	got := searchRange(nums, target)
	want := []int{3, 4}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
	got = searchRange(nums, 10)
	want = []int{5, 5}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
	got = searchRange([]int{}, 10)
	want = []int{-1, -1}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
	got = searchRange([]int{2, 2}, 3)
	want = []int{-1, -1}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
