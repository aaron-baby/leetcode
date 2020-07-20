package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	var r [][]int
	if root == nil {
		return r
	}
	// queue to store level elements
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	startFromLeft := true
	for len(queue) > 0 {
		l := len(queue)
		row := make([]int, l)
		for i := 0; i < l; i++ {
			node := queue[i]
			// push next level elements
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			if startFromLeft {
				row[i] = node.Val
			} else {
				row[l-1-i] = node.Val
			}
		}
		r = append(r, row)
		queue = queue[l:]
		startFromLeft = !startFromLeft
	}

	return r
}
