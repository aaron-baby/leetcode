package main

import (
	"reflect"
	"testing"
)

func Test_maximumNum(t *testing.T) {
	type args struct {
		ts []string
		n  int
	}
	tests := []struct {
		name string
		args args
		want Result
	}{
		{"example", args{[]string{"2018-01-01 01:01:01", "2018-01-01 01:01:02", "2018-01-01 01:01:05", "2018-01-01 01:02:05"}, 2}, Result{"2018-01-01 01:01:01", 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumNum(tt.args.ts, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maximumNum() = %v, want %v", got, tt.want)
			}
		})
	}
}
