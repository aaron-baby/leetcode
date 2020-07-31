package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		Val: val,
	}
}

func (t *TreeNode) String() {
	//h := height(t)

}

func SameTree(a *TreeNode, b *TreeNode) bool {
	// 1. both empty
	if a == nil && b == nil {
		return true
	}

	// 2. both non-empty -> compare them
	if a != nil && b != nil {
		if a.Val != b.Val {
			return false
		}
		return SameTree(a.Left, b.Left) && SameTree(a.Right, b.Right)
	}

	// 3. one empty, one not -> false
	return false
}

func insertNode(s []interface{}, root *TreeNode, i int, l int) *TreeNode {
	if i < l && s[i] != nil {
		root = &TreeNode{
			Val: s[i].(int),
		}
		root.Left = insertNode(s, root.Left, 2*i+1, l)
		root.Right = insertNode(s, root.Right, 2*i+2, l)

	}

	return root
}
