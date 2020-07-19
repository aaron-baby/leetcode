package sort

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestMerge(t *testing.T) {
	input := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	got := merge(input)
	want := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
	input = [][]int{{1, 4}, {4, 5}}
	got = merge(input)
	want = [][]int{{1, 5}}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
	got = merge([][]int{})
	want = [][]int{}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
