package dynamic_programming

import "testing"

func TestNumDecodings(t *testing.T) {
	s := "12"
	got := numDecodings(s)
	want := 2
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
