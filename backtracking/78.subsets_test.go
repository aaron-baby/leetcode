package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsets(t *testing.T) {
	nums := []int{1, 2, 3}
	got := subsets(nums)
	want := [][]int{
		{3},
		{1},
		{2},
		{1, 2, 3},
		{1, 3},
		{2, 3},
		{1, 2},
		{},
	}
	a := assert.New(t)
	a.ElementsMatch(got, want)
}
