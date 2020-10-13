package linked_list

import "testing"

func TestMiddleNode(t *testing.T) {
	l1 := &ListNode{Val: 1}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 3}
	l4 := &ListNode{Val: 4}
	l5 := &ListNode{Val: 5}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l5
	got := middleNode(l1)
	want := l3
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
