package dynamic_programming

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	got := longestPalindrome(s)
	want := "bab"
	if got != want {
		t.Errorf("string %s: got %v want %v", s, got, want)
	}

	s = "cbbd"
	got = longestPalindrome(s)
	want = "bb"
	if got != want {
		t.Errorf("string %s: got %v want %v", s, got, want)
	}
	s = "aaaa"
	got = longestPalindrome(s)
	want = "aaaa"
	if got != want {
		t.Errorf("string %s: got %v want %v", s, got, want)
	}
}
