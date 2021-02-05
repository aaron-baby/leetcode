package depth_first_search

import (
	"gitlab.com/aaw/leetcode/tree"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	type args struct {
		root      *tree.TreeNode
		targetSum int
	}
	var r *tree.TreeNode
	s := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, nil, 1}
	r = tree.InsertNode(s, r, 0, len(s))
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{root: r, targetSum: 22}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.args.root, tt.args.targetSum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
