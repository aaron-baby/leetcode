package hash_table

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestFindWords(t *testing.T) {
	input := []string{"Hello", "Alaska", "Dad", "Peace"}
	got := findWords(input)
	want := []string{"Alaska", "Dad"}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
