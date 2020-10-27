package dynamic_programming

import "testing"

func TestClimbStairs(t *testing.T) {
	n, want := 2, 2
	got := climbStairs(n)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	n, want = 3, 3
	got = climbStairs(n)
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
