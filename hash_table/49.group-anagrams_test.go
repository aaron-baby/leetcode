package hash_table

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	got := groupAnagrams(strs)
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	a := assert.New(t)
	a.ElementsMatch(got, want)
}
