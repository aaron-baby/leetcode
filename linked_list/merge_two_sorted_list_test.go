package linked_list

import "testing"

func TestMergeTwoSortedList(t *testing.T) {
	l1 := sliceToLinkedList([]int{1, 2, 4})
	l2 := sliceToLinkedList([]int{1, 3, 4})
	want := sliceToLinkedList([]int{1, 1, 2, 3, 4, 4})
	got := mergeTwoLists(l1, l2)
	if !CompareLinkedList(got, want) {
		t.Errorf("got %s want %s given", got, want)
	}
}
