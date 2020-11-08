package backtracking

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	candidates := []int{2, 3, 6, 7}
	target := 7
	got := combinationSum(candidates, target)
	want := [][]int{
		{7},
		{2, 2, 3},
	}
	assert := assert.New(t)
	assert.ElementsMatch(got, want)

	fmt.Println("=====")
	candidates = []int{2, 3, 5}
	target = 8
	got = combinationSum(candidates, target)
	want = [][]int{
		{2, 2, 2, 2},
		{2, 3, 3},
		{3, 5},
	}
	assert.ElementsMatch(got, want)
}
