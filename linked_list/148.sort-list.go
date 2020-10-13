package linked_list

// 148. Sort List
// https://leetcode.com/problems/sort-list/
func sortList(head *ListNode) *ListNode {
	// empty list or last node
	if head == nil || head.Next == nil {
		return head
	}
	mid := middleNodeBreak(head)
	left := sortList(head)
	right := sortList(mid)
	return mergeTwoLists(left, right)
}

func middleNodeBreak(head *ListNode) *ListNode {
	slow, slowPrev, fast := head, head, head
	for fast != nil && fast.Next != nil {
		slowPrev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	slowPrev.Next = nil
	return slow
}
