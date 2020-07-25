package linked_list

import (
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	n := 2
	removeNthFromEnd(l1, n)
	if l3.Next != l5 {
		t.Errorf("%dth node from the end didn't remove correctly\n%s", n, l1.String())
	}

	node := &ListNode{Val: 1}
	h := removeNthFromEnd(node, 1)
	if h != nil {
		t.Errorf("%dth node from the end didn't remove from %s correctly", n, node.String())
	}
}
