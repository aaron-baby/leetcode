package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubsetsWithDup(t *testing.T) {
	nums := []int{1, 2, 2}
	got := subsetsWithDup(nums)
	want := [][]int{
		{2},
		{1},
		{1, 2, 2},
		{2, 2},
		{1, 2},
		{},
	}
	a := assert.New(t)
	a.ElementsMatch(got, want)
}
