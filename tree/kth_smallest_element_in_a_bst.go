package tree

import "github.com/golang-collections/collections/stack"

// 230. Kth Smallest Element in a BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
func kthSmallest(root *TreeNode, k int) int {
	if root == nil || k == 0 {
		return -1
	}
	st := stack.New()
	for {
		// push node to stack until to smallest element
		for root != nil {
			st.Push(root)
			root = root.Left
		}
		// pop smallest element in stack
		root = st.Pop().(*TreeNode)
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}
