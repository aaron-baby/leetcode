package tree

import "testing"

func TestIsSymmetric(t *testing.T) {
	var r *TreeNode
	s := []interface{}{1, 2, 2, 3, 4, 4, 3}
	r = insertNode(s, r, 0, len(s))
	got := isSymmetric(r)
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
