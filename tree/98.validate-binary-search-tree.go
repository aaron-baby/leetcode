package tree

func isValidBST(root *TreeNode) bool {
	seq := inorderTraversal(root)
	for i := 1; i < len(seq); i++ {
		if seq[i-1] >= seq[i] {
			return false
		}
	}
	return true
}
