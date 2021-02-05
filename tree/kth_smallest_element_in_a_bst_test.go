package tree

import "testing"

func TestKthSmallest(t *testing.T) {
	var root, r1 *TreeNode
	s := []interface{}{3, 1, 4, nil, 2}
	root = InsertNode(s, root, 0, len(s))
	got := kthSmallest(root, 1)
	want := 1
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

	s = []interface{}{5, 3, 6, 2, 4, nil, nil, 1}
	r1 = InsertNode(s, r1, 0, len(s))
	got = kthSmallest(r1, 3)
	want = 3
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
