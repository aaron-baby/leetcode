package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/tree"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func GenerateTrees(n int) []*tree.TreeNode {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i + 1
	}
	return getPermutation(s)
}

// given a slice, pick a root, generate left and right subtree, return root node
func getPermutation(s []int) []*tree.TreeNode {
	if 0 == len(s) {
		return []*tree.TreeNode{nil}
	}
	if 1 == len(s) {
		r := &tree.TreeNode{Val: s[0]}
		return []*tree.TreeNode{r}
	}
	var rst []*tree.TreeNode
	for i := 0; i < len(s); i++ {
		left := getPermutation(s[:i])
		right := getPermutation(s[i+1:])
		for _, lr := range left {
			for _, rr := range right {
				rst = append(rst, &tree.TreeNode{Val: s[i], Left: lr, Right: rr})
			}
		}
	}
	return rst
}
