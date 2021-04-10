package dynamic_programming

import "testing"

func TestWordBreak(t *testing.T) {
	testCases := map[string][]string{
		"leetcode":      []string{"leet", "code"},
		"applepenapple": []string{"apple", "pen"},
	}
	wants := []bool{true, true}
	i := 0
	for s, wordDict := range testCases {
		got := wordBreak(s, wordDict)
		if got != wants[i] {
			t.Errorf("string %s: got %t want %t", s, got, wants[i])
		}
		i++
	}
}
