package main

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	m := map[string]int{
		"abcabcbb": 3,
		"bbbbb":    1,
		"pwwkew":   3,
	}

	for s, want := range m {
		got := lengthOfLongestSubstring(s)
		if got != want {
			t.Errorf("for string %s got %v want %v", s, got, want)
		}

	}
}
