package tree

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var q1, q2 []*TreeNode
	q1 = append(q1, root.Left)
	q2 = append(q2, root.Right)
	for len(q1) > 0 && len(q2) > 0 {
		var node1, node2 *TreeNode
		node1, q1 = q1[len(q1)-1], q1[:len(q1)-1]
		node2, q2 = q2[len(q2)-1], q2[:len(q2)-1]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil {
			return false
		}
		if node1.Val != node2.Val {
			return false
		}
		q1 = append(q1, node1.Left)
		q1 = append(q1, node1.Right)
		q2 = append(q2, node2.Right)
		q2 = append(q2, node2.Left)
	}
	return true
}
