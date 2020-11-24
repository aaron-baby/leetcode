package stack

import "testing"

func TestSimplifyPath(t *testing.T) {
	path := []string{"/home/", "/../", "/home//foo/", "/a/./b/../../c/"}
	wants := []string{"/home", "/", "/home/foo", "/c"}
	var got string
	for i, v := range path {
		got = simplifyPath(v)
		if got != wants[i] {
			t.Errorf("got %v want %v", got, wants[i])
		}
	}
}
