package stack

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	got := dailyTemperatures(T)
	want := []int{1, 1, 4, 2, 1, 1, 0, 0}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
