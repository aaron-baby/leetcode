package union_find

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{"Example 1", args{[][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}}}, [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}}},
		{"Testcase 13", args{[][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}, {'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}}, [][]byte{{'X', 'O', 'X', 'O', 'X', 'O'}, {'O', 'X', 'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X', 'X', 'O'}, {'O', 'X', 'O', 'X', 'O', 'X'}}},
		{"Testcase 25", args{[][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}}, [][]byte{{'O', 'O', 'O'}, {'O', 'O', 'O'}, {'O', 'O', 'O'}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if solve(tt.args.board); !cmp.Equal(tt.args.board, tt.want) {
				t.Errorf("solve() = %v, want %v", tt.args.board, tt.want)
			}
		})
	}
}
