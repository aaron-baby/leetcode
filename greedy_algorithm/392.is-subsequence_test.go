package greedy_algorithm

import "testing"

func Test_isSubsequence(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Example 1", args{"abc", "ahbgdc"}, true},
		{"Example 2", args{"axc", "ahbgdc"}, false},
		{"Example 3", args{"ahbgdc", "ahbgdc"}, true},
		{"Example 4", args{"ahbgdc", "ahbgde"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSubsequence(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("isSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
