package linked_list

import "testing"

func TestReorderList(t *testing.T) {
	s := []int{1, 2, 3, 4}
	h := sliceToLinkedList(s)
	reorderList(h)
	want := sliceToLinkedList([]int{1, 4, 2, 3})
	if !CompareLinkedList(h, want) {
		t.Errorf("got %s want %s given", h, want)
	}

	got := sliceToLinkedList([]int{1, 2, 3, 4, 5})
	reorderList(got)
	want = sliceToLinkedList([]int{1, 5, 2, 4, 3})
	if !CompareLinkedList(got, want) {
		t.Errorf("got %s want %s given", h, want)
	}
}
