package main

import (
	"testing"
	"github.com/google/go-cmp/cmp"
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

func insertNode(s []interface{}, root *TreeNode, i int, l int) *TreeNode {
	if i < l && s[i] != nil {
		root = &TreeNode{Val: (s[i].(int))}
		root.Left = insertNode(s, root.Left, 2*i+1, l)
		root.Right = insertNode(s, root.Right, 2*i+2, l)

	}

	return root
}
