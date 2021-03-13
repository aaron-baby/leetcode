package breadth_first_search

import (
	"gitlab.com/aaw/leetcode/tree"
	"reflect"
	"testing"
)

func Test_rightSideView(t *testing.T) {
	type args struct {
		root *tree.TreeNode
	}
	l := &tree.TreeNode{Val: 2, Right: &tree.TreeNode{Val: 5}}
	r := &tree.TreeNode{Val: 3, Right: &tree.TreeNode{Val: 4}}
	tr := &tree.TreeNode{Val: 1, Left: l, Right: r}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Example 1", args{tr}, []int{1, 3, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rightSideView(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView() = %v, want %v", got, tt.want)
			}
		})
	}
}
