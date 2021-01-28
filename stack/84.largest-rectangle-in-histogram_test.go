package stack

import "testing"

func Test_largestRectangleArea(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]int{2, 1, 5, 6, 2, 3}}, 10},
		{"Example+1", args{[]int{2, 1, 6, 5, 6, 2, 3}}, 15},
		{"decrease", args{[]int{7, 6, 5, 4, 3, 2}}, 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestRectangleArea(tt.args.heights); got != tt.want {
				t.Errorf("largestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
