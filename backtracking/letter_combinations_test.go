package backtracking

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLetterCombinations(t *testing.T) {
	input := "23"
	got := letterCombinations(input)
	want := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
