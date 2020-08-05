package linked_list

import "testing"

func TestMergeKLists(t *testing.T) {
	l1 := sliceToLinkedList([]int{1, 4, 5})
	l2 := sliceToLinkedList([]int{1, 3, 4})
	l3 := sliceToLinkedList([]int{2, 6})
	got := mergeKLists([]*ListNode{l1, l2, l3})
	want := sliceToLinkedList([]int{1, 1, 2, 3, 4, 4, 5, 6})
	if !CompareLinkedList(got, want) {
		t.Errorf("got %s want %s given", got, want)
	}
}
