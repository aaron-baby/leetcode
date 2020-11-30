package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	got := spiralOrder(matrix)
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
