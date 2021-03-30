package math

import "testing"

func Test_myAtoi(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1:", args{"42"}, 42},
		{"Example 2:", args{"   -42"}, -42},
		{"Example 3:", args{"4193 with words"}, 4193},
		{"Example 4:", args{"words and 987"}, 0},
		{"Example 5:", args{"-91283472332"}, -2147483648},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.args.s); got != tt.want {
				t.Errorf("myAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
