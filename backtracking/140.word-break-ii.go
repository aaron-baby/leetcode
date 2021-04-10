package backtracking

import (
	"strings"
)

func wordBreakii(s string, wordDict []string) []string {
	var res = []string{}
	c := []string{}
	findComb(s, wordDict, c, &res)
	return res
}

func findComb(remain string, wordDict []string, c []string, res *[]string) {
	if remain == "" {
		*res = append(*res, strings.Join(c, " "))
		return
	}
	for _, w := range wordDict {
		if len(w) > len(remain) || w != remain[0:len(w)] {
			continue
		}
		c = append(c, w)
		findComb(remain[len(w):], wordDict, c, res)
		c = c[:len(c)-1]
	}
}
