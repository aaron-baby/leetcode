package backtracking

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_wordBreakii(t *testing.T) {
	type args struct {
		s        string
		wordDict []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example 1:", args{"catsanddog", []string{"cat", "cats", "and", "sand", "dog"}}, []string{"cats and dog", "cat sand dog"}},
	}
	as := assert.New(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wordBreakii(tt.args.s, tt.args.wordDict); !as.ElementsMatch(got, tt.want) {
				t.Errorf("wordBreakii() = %v, want %v", got, tt.want)
			}
		})
	}
}
