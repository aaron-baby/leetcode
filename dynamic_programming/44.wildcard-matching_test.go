package dynamic_programming

import "testing"

func TestIsMatch44(t *testing.T) {
	s, p := "aa", "a"
	got := isMatch44(s, p)
	want := false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
	s, p = "aa", "*"
	got = isMatch44(s, p)
	want = true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	s, p = "cb", "?a"
	got = isMatch44(s, p)
	want = false
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}

	s, p = "adceb", "*a*b"
	got = isMatch44(s, p)
	want = true
	if got != want {
		t.Errorf("got %v want %v given", got, want)
	}
}
