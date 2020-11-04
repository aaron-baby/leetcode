package two_pointers

import "testing"

func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	got := trap(height)
	want := 6
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
