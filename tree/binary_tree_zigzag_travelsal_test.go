package tree

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestBinaryTreeZigzagTravelsal(t *testing.T) {
	var r *TreeNode
	s := []interface{}{3, 9, 20, nil, nil, 15, 7}
	r = insertNode(s, r, 0, len(s))
	got := zigzagLevelOrder(r)
	want := [][]int{
		{3},
		{20, 9},
		{15, 7},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
