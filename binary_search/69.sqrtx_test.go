package binary_search

import "testing"

func TestMySqrt(t *testing.T) {
	x := 4
	got := mySqrt(x)
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
