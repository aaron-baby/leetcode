package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestRotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	want := [][]int{
		{7, 4, 1},
		{8, 5, 2},
		{9, 6, 3},
	}
	if !cmp.Equal(matrix, want) {
		t.Errorf("got %v want %v given", matrix, want)
	}
}
