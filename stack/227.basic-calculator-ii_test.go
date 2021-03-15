package stack

import "testing"

func Test_calculate_ii(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"3+2*2"}, 7},
		{"Example 2", args{" 3/2 "}, 1},
		{"Example 3", args{" 3+5 / 2 "}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate_ii(tt.args.s); got != tt.want {
				t.Errorf("calculate_ii() = %v, want %v", got, tt.want)
			}
		})
	}
}
