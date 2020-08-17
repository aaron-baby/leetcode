package bit_manipulation

import "testing"

func TestNumSteps(t *testing.T) {
	ss := []string{"1101", "10", "1"}
	wants := []int{6, 1, 0}
	for i, s := range ss {
		got := numSteps(s)
		if got != wants[i] {
			t.Errorf("got %v want %v", got, wants[i])
		}
	}
}
