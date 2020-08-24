package tree

func findMode(root *TreeNode) []int {
	var res []int
	max, cnt := 0, 1
	var pre *TreeNode
	inOrder(root, &pre, &cnt, &max, &res)
	return res
}

func inOrder(node *TreeNode, pre **TreeNode, cnt *int, max *int, res *[]int) {
	if node == nil {
		return
	}
	inOrder(node.Left, pre, cnt, max, res)
	q := &pre
	n := **q
	if n != nil && n.Val == node.Val {
		*cnt++
	} else {
		*cnt = 1
	}
	output := *res
	if *cnt >= *max {
		if *cnt > *max {
			output = []int{}
		}
		*res = append(output, node.Val)
		*max = *cnt
	}
	**q = node
	inOrder(node.Right, pre, cnt, max, res)
}
