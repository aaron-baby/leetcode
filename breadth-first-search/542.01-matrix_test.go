package breadth_first_search

import (
	"reflect"
	"testing"
)

func Test_updateMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"Example 1", args{matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{"Example 2", args{matrix: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := updateMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("updateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
