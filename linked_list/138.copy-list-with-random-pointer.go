package linked_list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var res *Node
	if head == nil {
		return res
	}
	res = &Node{Val: head.Val}
	node := res
	cur := head.Next
	m := make(map[*Node]*Node)
	m[head] = res
	for cur != nil {
		t := &Node{Val: cur.Val}
		node.Next = t
		m[cur] = t
		node = node.Next
		cur = cur.Next
	}
	// set random
	node = res
	cur = head
	for cur != nil {
		node.Random = m[cur.Random]
		node = node.Next
		cur = cur.Next
	}
	return res
}
