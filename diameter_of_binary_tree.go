package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight, rHeight := height(root.Left), height(root.Right)
	lDiameter, rDiameter := diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)

	return max(lHeight+rHeight, max(lDiameter, rDiameter))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}

	return max(height(node.Left), height(node.Right)) + 1
}
