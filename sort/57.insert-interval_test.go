package sort

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInsert(t *testing.T) {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	got := insert(intervals, newInterval)
	want := [][]int{{1, 5}, {6, 9}}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
