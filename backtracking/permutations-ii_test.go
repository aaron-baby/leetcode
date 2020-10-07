package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	nums := []int{1, 1, 2}
	results := permuteUnique(nums)
	a := assert.New(t)
	a.ElementsMatch(results, [][]int{
		{1, 1, 2},
		{1, 2, 1},
		{2, 1, 1},
	})

	nums = []int{2, 2, 1, 1}
	results = permuteUnique(nums)
	a.ElementsMatch(results, [][]int{
		{1, 1, 2, 2},
		{1, 2, 1, 2},
		{1, 2, 2, 1},
		{2, 1, 1, 2},
		{2, 1, 2, 1},
		{2, 2, 1, 1},
	})
}
