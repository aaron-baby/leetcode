package main

import "testing"

func TestValidParentheses(t *testing.T) {
	testCases := map[string]bool{
		"()":     true,
		"()[]{}": true,
		"(]":     false,
		"([)]":   false,
		"{[]}":   true,
		"]":      false,
		"[])":      false,
	}
	for s, want := range testCases {
		got := isValid(s)
		if got != want {
			t.Errorf("string %s: got %t want %t", s, got, want)
		}
	}
}
