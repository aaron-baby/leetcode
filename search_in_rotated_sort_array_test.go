package main

import "testing"

func TestSearchInRotatedSortArray(t *testing.T) {
	nums := []int{4, 5, 6, 7, 0, 1, 2}
	target := 0
	got := search(nums, target)
	expected := 4
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
	expected = -1
	got = search(nums, 3)
	if got != expected {
		t.Errorf("expected '%d' but got '%d'", expected, got)
	}
}
