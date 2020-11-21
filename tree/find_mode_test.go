package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMode(t *testing.T) {
	r := &TreeNode{
		Val:  2,
		Left: NewTreeNode(2),
	}
	root := &TreeNode{
		Val:   1,
		Right: r,
	}
	var got, want []int
	got = findMode(root)
	want = []int{2}
	a := assert.New(t)
	a.ElementsMatch(got, want)

	var r1, r2 *TreeNode
	s := []interface{}{1, nil, 2}
	r1 = insertNode(s, r1, 0, len(s))
	got = findMode(r1)
	want = []int{2, 1}
	a.ElementsMatch(got, want)

	s = []interface{}{2, 1}
	r2 = insertNode(s, r2, 0, len(s))
	got = findMode(r2)
	want = []int{2, 1}
	a.ElementsMatch(got, want)
}
