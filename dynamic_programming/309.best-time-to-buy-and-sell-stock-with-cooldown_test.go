package dynamic_programming

import "testing"

func Test_maxProfitCooldown(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{[]int{1, 2, 3, 0, 2}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitCooldown(tt.args.prices); got != tt.want {
				t.Errorf("maxProfitCooldown() = %v, want %v", got, tt.want)
			}
		})
	}
}
