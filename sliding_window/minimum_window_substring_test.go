package sliding_window

import "testing"

func TestMinWindow(t *testing.T) {
	s := "ADOBECODEBANC"
	subset := "ABC"
	want := "BANC"
	got := minWindow(s, subset)
	if got != want {
		t.Errorf("for string %s, subset %s got %v want %v", s, subset, got, want)
	}
	got = minWindow("a", "aa")
	if got != "" {
		t.Errorf("for string %s, subset %s got %v want %v", s, subset, got, want)
	}
}
