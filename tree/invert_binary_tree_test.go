package tree

import "testing"

func TestInvertTree(t *testing.T) {
	var r, wantTree *TreeNode
	s := []interface{}{4, 2, 7, 1, 3, 6, 9}
	r = InsertNode(s, r, 0, len(s))
	invertTree(r)

	want := []interface{}{4, 7, 2, 9, 6, 3, 1}
	wantTree = InsertNode(want, wantTree, 0, len(want))
	if !SameTree(r, wantTree) {
		t.Errorf("got %v want %v given", r, wantTree)
	}
}
