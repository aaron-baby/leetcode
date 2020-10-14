package backtracking

import "testing"

func TestNumSquarefulPerms(t *testing.T) {
	A := []int{1, 17, 8}
	got := numSquarefulPerms(A)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	A = []int{2, 2, 2}
	got = numSquarefulPerms(A)
	want = 1
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
