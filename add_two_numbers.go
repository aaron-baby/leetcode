package main

import "gitlab.com/aaw/leetcode/ds"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ds.ListNode, l2 *ds.ListNode) *ds.ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}
	var carry, sum int
	var result, temp, prev *ds.ListNode
	for l1 != nil || l2 != nil {
		sum = carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// quotient
		carry = sum / 10
		// reminder
		sum = sum % 10
		temp = &ds.ListNode{Val: sum}
		if result == nil {
			result = temp
		} else {
			prev.Next = temp
		}
		prev = temp
	}
	if carry > 0 {
		temp.Next = &ds.ListNode{Val: carry}
	}

	return result
}
