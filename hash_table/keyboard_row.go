package hash_table

import (
	"strings"
)

func findWords(words []string) []string {
	var ans []string
	for _, word := range words {
		if inSameRow(word) {
			ans = append(ans, word)
		}
	}
	return ans
}

func inSameRow(word string) bool {
	var letterRow = make(map[rune]int)
	for _, l := range []rune{'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p'} {
		letterRow[l] = 1
	}
	for _, l := range []rune{'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l'} {
		letterRow[l] = 2
	}
	for _, l := range []rune{'z', 'x', 'c', 'v', 'b', 'n', 'm'} {
		letterRow[l] = 3
	}
	var rn int
	for i, r := range strings.ToLower(word) {
		if i == 0 {
			rn = letterRow[r]
			continue
		}
		if letterRow[r] != rn {
			return false
		}
	}
	return true
}
