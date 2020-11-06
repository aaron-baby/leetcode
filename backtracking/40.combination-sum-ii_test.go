package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	got := combinationSum2(candidates, target)
	want := [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}
	a := assert.New(t)
	a.ElementsMatch(got, want)
}
