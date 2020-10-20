package dynamic_programming

import "testing"

func TestIsMatch(t *testing.T) {
	s, p := "aa", "a"
	got := isMatch(s, p)
	want := false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
