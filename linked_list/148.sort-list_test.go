package linked_list

import "testing"

func TestSortList(t *testing.T) {
	l4 := &ListNode{Val: 3}
	l3 := &ListNode{Val: 1, Next: l4}
	l2 := &ListNode{Val: 2, Next: l3}
	l1 := &ListNode{Val: 4, Next: l2}
	got := sortList(l1)
	want := sliceToLinkedList([]int{1, 2, 3, 4})
	if !CompareLinkedList(got, want) {
		t.Errorf("got %s want %s given", got, want)
	}
}
