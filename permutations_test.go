package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPermute(t *testing.T) {
	res := permute([]int{1, 2, 3})
	assert := assert.New(t)
	assert.ElementsMatch(res, [][]int{
		{1, 2, 3},
		{1, 3, 2},
		{2, 1, 3},
		{2, 3, 1},
		{3, 1, 2},
		{3, 2, 1},
	})
}
