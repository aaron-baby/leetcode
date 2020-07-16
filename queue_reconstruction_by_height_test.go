package main

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestReconstructQueue(t *testing.T) {
	people := [][]int{
		{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
	}
	want := [][]int{
		{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1},
	}
	got := reconstructQueue(people)
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
