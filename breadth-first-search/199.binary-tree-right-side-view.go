package breadth_first_search

import "gitlab.com/aaw/leetcode/tree"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *tree.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	var q []*tree.TreeNode
	q = append(q, root)
	for len(q) > 0 {
		size := len(q)
		for size > 0 {
			size--
			cur := q[0]
			q = q[1:]
			if size == 0 {
				res = append(res, cur.Val)
			}
			if cur.Left != nil {
				q = append(q, cur.Left)
			}
			if cur.Right != nil {
				q = append(q, cur.Right)
			}
		}
	}
	return res
}
