package tree

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// p and q are different and both values will exist in the binary tree.
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	// if p and q are not in root node, it must exist in left or right subtree
	nodeInLeft := lowestCommonAncestor(root.Left, p, q)
	nodeInRight := lowestCommonAncestor(root.Right, p, q)
	// nodes in different subtrees
	if nodeInLeft != nil && nodeInRight != nil {
		return root
	}
	// both p and q are in left subtree
	if nodeInLeft != nil {
		return nodeInLeft
	}
	return nodeInRight
}
