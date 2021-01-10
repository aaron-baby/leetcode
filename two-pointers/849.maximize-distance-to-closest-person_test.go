package two_pointers

import "testing"

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]int{1, 0, 0, 0, 1, 0, 1}}, 2},
		{"Example 2", args{[]int{1, 0, 0, 0}}, 3},
		{"Example 3", args{[]int{0, 1}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDistToClosest(tt.args.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
