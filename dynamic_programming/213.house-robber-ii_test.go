package dynamic_programming

import "testing"

func Test_robii(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{2, 3, 2}}, 3},
		{"Example 2", args{[]int{1, 2, 3, 1}}, 4},
		{"Example 3", args{[]int{0}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robii(tt.args.nums); got != tt.want {
				t.Errorf("robii() = %v, want %v", got, tt.want)
			}
		})
	}
}
