package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThreeSum(t *testing.T) {
	nums := []int{-1, 0, 1, 2, -1, -4}
	got := threeSum(nums)
	want := [][]int{
		{-1, -1, 2},
		{-1, 0, 1},
	}
	a := assert.New(t)
	a.ElementsMatch(want, got)

	nums = []int{1, -1, -1, 0}
	got = threeSum(nums)
	want = [][]int{
		{-1, 0, 1},
	}
	a.ElementsMatch(want, got)
}
