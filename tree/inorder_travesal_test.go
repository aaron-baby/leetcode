package tree

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestInorderTravelsal(t *testing.T) {
	root := NewTreeNode(1)
	rnode := NewTreeNode(2)
	root.Right = rnode
	rnode.Left = NewTreeNode(3)

	got := inorderTraversal(root)
	want := []int{1, 3, 2}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
