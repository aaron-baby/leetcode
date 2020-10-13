package greedy_algorithm

import (
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	S := "ababcbacadefegdehijhklij"
	got := partitionLabels(S)
	want := []int{9, 7, 8}
	if !cmp.Equal(got, want) {
		t.Errorf("got %v want %v given", got, want)
	}
}
