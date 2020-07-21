package linked_list

import (
	"testing"
)

func TestLinkedListCycle(t *testing.T) {
	l1 := &ListNode{Val: 3}
	l2 := &ListNode{Val: 2}
	l3 := &ListNode{Val: 0}
	l4 := &ListNode{Val: -4}
	l1.Next = l2
	l2.Next = l3
	l3.Next = l4
	l4.Next = l2
	want := true
	got := hasCycle(l1)
	if got != want {
		t.Errorf("given list node l1:\n%s\ngot %t want %t", l1.String(), got, want)
	}
}
