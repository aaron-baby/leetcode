package array

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	got := productExceptSelf(nums)
	want := []int{24, 12, 8, 6}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
