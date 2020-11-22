package backtracking

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGrayCode(t *testing.T) {
	n := 2
	got := grayCode(n)
	want := []int{0, 1, 3, 2}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
