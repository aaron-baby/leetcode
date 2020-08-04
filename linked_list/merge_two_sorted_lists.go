package linked_list

// 21. Merge Two Sorted Lists
// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var prehead = &ListNode{Val: -1}
	cur := prehead
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	cur.Next = l2
	if l1 != nil {
		cur.Next = l1
	}

	return prehead.Next
}
