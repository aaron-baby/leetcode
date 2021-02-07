package depth_first_search

import (
	"gitlab.com/aaw/leetcode/tree"
	"reflect"
	"testing"
)

func Test_pathSum(t *testing.T) {
	type args struct {
		root      *tree.TreeNode
		targetSum int
	}
	var r *tree.TreeNode
	s := []interface{}{5, 4, 8, 11, nil, 13, 4, 7, 2, nil, nil, 5, 1}
	r = tree.InsertNode(s, r, 0, len(s))

	root := tree.NewTreeNode(5)
	root.Left = tree.NewTreeNode(4)
	root.Left.Left = tree.NewTreeNode(11)
	root.Left.Left.Left = tree.NewTreeNode(7)
	root.Left.Left.Right = tree.NewTreeNode(2)

	root.Right = tree.NewTreeNode(8)
	root.Right.Left = tree.NewTreeNode(13)
	root.Right.Right = tree.NewTreeNode(4)
	root.Right.Right.Left = tree.NewTreeNode(5)
	root.Right.Right.Right = tree.NewTreeNode(1)
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1:", args{root: root, targetSum: 22}, [][]int{{5,4,11,2},{5,8,4,5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := pathSum(tt.args.root, tt.args.targetSum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
