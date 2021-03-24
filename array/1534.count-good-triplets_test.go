package array

import "testing"

func Test_countGoodTriplets(t *testing.T) {
	type args struct {
		arr []int
		a   int
		b   int
		c   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1:", args{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("countGoodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
