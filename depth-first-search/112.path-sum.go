package depth_first_search

import "gitlab.com/aaw/leetcode/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *tree.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	var qn []*tree.TreeNode
	var qv []int
	qn = append(qn, root)
	qv = append(qv, root.Val)
	for len(qn) > 0 {
		// front
		n := qn[0]
		qn = qn[1:]
		v := qv[0]
		qv = qv[1:]
		if n.Left == nil && n.Right == nil {
			if v == targetSum {
				return true
			}
			continue
		}
		if n.Left != nil {
			qn = append(qn, n.Left)
			qv = append(qv, v+n.Left.Val)
		}
		if n.Right != nil {
			qn = append(qn, n.Right)
			qv = append(qv, v+n.Right.Val)
		}
	}
	return false
}
