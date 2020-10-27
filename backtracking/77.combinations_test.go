package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCombine(t *testing.T) {
	n, k := 4, 2
	got := combine(n, k)
	want := [][]int{
		{2, 4},
		{3, 4},
		{2, 3},
		{1, 2},
		{1, 3},
		{1, 4},
	}
	a := assert.New(t)
	a.ElementsMatch(got, want)
}
