package tree

import "github.com/golang-collections/collections/stack"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
	curr := root
	stack := stack.New()
	res := []int{}
	for curr != nil || stack.Len() > 0 {
		for curr != nil {
			stack.Push(curr)
			curr = curr.Left
		}
		curr = stack.Pop().(*TreeNode)
		res = append(res, curr.Val)
		curr = curr.Right
	}

	return res
}
