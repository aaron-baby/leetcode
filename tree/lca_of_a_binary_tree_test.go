package tree

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	var root *TreeNode
	s := []interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}
	root = insertNode(s, root, 0, len(s))
	got := lowestCommonAncestor(root, NewTreeNode(5), NewTreeNode(1)).Val
	want := 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
