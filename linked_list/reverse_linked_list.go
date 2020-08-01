package linked_list

func reverseList(head *ListNode) *ListNode {
	// Initialize prev and current pointers
	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		// point to prev
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
