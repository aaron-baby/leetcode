package tree

import (
	"gitlab.com/aaw/leetcode/lib"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var nodeHeight = make(map[*TreeNode]int)

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lHeight, rHeight := height(root.Left), height(root.Right)
	lDiameter, rDiameter := diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)

	return lib.Max(lHeight+rHeight, lib.Max(lDiameter, rDiameter))
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	_, ok := nodeHeight[node]
	if !ok {
		nodeHeight[node] = lib.Max(height(node.Left), height(node.Right)) + 1
	}

	return nodeHeight[node]
}
