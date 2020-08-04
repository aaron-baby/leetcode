package backtracking

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestGenerateParentheses(t *testing.T) {
	n := 3
	got := generateParenthesis(n)
	want := []string{
		"((()))",
		"(()())",
		"(())()",
		"()(())",
		"()()()",
	}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
