package stack

import (
	"testing"
)

func TestLongestValidParentheses(t *testing.T) {
	s := "(()"
	want := 2
	got := longestValidParentheses(s)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}

	s = "())"
	want = 2
	got = longestValidParentheses(s)
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
