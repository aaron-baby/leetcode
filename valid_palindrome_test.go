package main

import "testing"

func TestValidPalindrome(t *testing.T) {
	cases := map[string]bool{
		"abc":  false,
		"abca": true,
		"a":    true,
	}

	for s, want := range cases {
		got := validPalindrome(s)
		if got != want {
			t.Errorf("string %s: got %t want %t", s, got, want)
		}
	}
}
