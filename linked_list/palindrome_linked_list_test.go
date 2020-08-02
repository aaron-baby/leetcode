package linked_list

import "testing"

func TestIsPalindrome(t *testing.T) {
	l := sliceToLinkedList([]int{1, 2})
	got := isPalindrome(l)
	want := false
	if got != want {
		t.Errorf("given linked list l:\n%s\ngot %t want %t", l.String(), got, want)
	}

	l = sliceToLinkedList([]int{1, 2, 5, 3, 5, 2, 1})
	got = isPalindrome(l)
	want = true
	if got != want {
		t.Errorf("given linked list l:\n%s\ngot %t want %t", l.String(), got, want)
	}
}
