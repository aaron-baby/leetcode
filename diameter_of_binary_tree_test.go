package main

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	root := NewTreeNode(1)
	lnode := NewTreeNode(2)
	lnode.Left = NewTreeNode(4)
	lnode.Right = NewTreeNode(5)
	root.Left = lnode
	root.Right = NewTreeNode(3)
	got := diameterOfBinaryTree(root)
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}
