package main

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	input := []int{1, 2, 0}
	want := 3
	got := firstMissingPositive(input)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
