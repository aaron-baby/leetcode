package depth_first_search

import "gitlab.com/aaw/leetcode/tree"

func pathSum(root *tree.TreeNode, targetSum int) [][]int {
	var path []int
	var paths [][]int
	findPaths(root, targetSum, path, &paths)
	return paths
}

func findPaths(node *tree.TreeNode, sum int, path []int, paths *[][]int) {
	if node == nil {
		return
	}
	path = append(path, node.Val)
	if node.Left == nil && node.Right == nil && sum == node.Val {
		p := make([]int, len(path))
		copy(p, path)
		*paths = append(*paths, p)
		return
	}
	findPaths(node.Left, sum-node.Val, path, paths)
	findPaths(node.Right, sum-node.Val, path, paths)
	path = path[:len(path)-1]
}
