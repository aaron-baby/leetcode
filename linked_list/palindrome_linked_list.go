package linked_list

// 234. Palindrome Linked List
// https://leetcode.com/problems/palindrome-linked-list/
func isPalindrome(head *ListNode) bool {
	// empty or one element
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre := head
	post := slow.Next
	for post.Next != nil {
		tmp := post.Next
		post.Next = tmp.Next
		tmp.Next = slow.Next
		slow.Next = tmp
	}
	for slow.Next != nil {
		slow = slow.Next
		if pre.Val != slow.Val {
			return false
		}
		pre = pre.Next
	}
	return true
}
