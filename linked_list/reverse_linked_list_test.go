package linked_list

import "testing"

func TestReverseList(t *testing.T) {
	head := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	got := reverseList(head)
	want := sliceToLinkedList([]int{5, 4, 3, 2, 1})
	if !CompareLinkedList(got, want) {
		t.Errorf("got %s want %s given", got, want)
	}
}
