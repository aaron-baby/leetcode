package math

import "testing"

func Test_maxPoints(t *testing.T) {
	type args struct {
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1:", args{points: [][]int{{1, 1}, {2, 2}, {3, 3}}}, 3},
		{"two dots:", args{points: [][]int{{1, 1}, {2, 3}}}, 2},
		{"tc 3:", args{points: [][]int{{-6, -1}, {3, 1}, {12, 3}}}, 3},
		{"tc 4:", args{points: [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPoints(tt.args.points); got != tt.want {
				t.Errorf("maxPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
