package linked_list

func reorderList(head *ListNode) {
	// empty or just have one node
	if head == nil || head.Next == nil {
		return
	}
	// 1) Find the middle point using tortoise and hare method
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 2) Split the linked list into two halves
	head1 := head
	head2 := slow.Next
	slow.Next = nil

	// 3) Reverse the second half
	head2 = reverseList(head2)

	// 4) alternate merge of first and second halves
	dummy := &ListNode{Val: 0}
	curr := dummy
	for head1 != nil || head2 != nil {
		if head1 != nil {
			curr.Next = head1
			curr = curr.Next
			head1 = head1.Next
		}
		if head2 != nil {
			curr.Next = head2
			curr = curr.Next
			head2 = head2.Next
		}
	}
	head = dummy.Next
}
