package linked_list

func insertionSortList(head *ListNode) *ListNode {
	dummyHead := &ListNode{}
	cur := head
	var pre, nextNode *ListNode
	for cur != nil {
		pre = dummyHead
		nextNode = pre.Next
		// finds the location in sorted list
		for nextNode != nil {
			if cur.Val < nextNode.Val {
				break
			}
			pre = nextNode
			nextNode = nextNode.Next
		}
		// cursor's next
		nextItem := cur.Next
		// insertion
		cur.Next = nextNode
		pre.Next = cur

		cur = nextItem
	}
	return dummyHead.Next
}
