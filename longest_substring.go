package main

func lengthOfLongestSubstring(s string) int {
	var i, maxLen int
	byteMap := make(map[byte]int)
	// j is right frame of sliding window
	for j := 0; j < len(s); j++ {
		lastIndex, ok := byteMap[s[j]]
		if ok {
			i = max(i, lastIndex+1)
		}
		if j-i+1 > maxLen {
			maxLen = j - i + 1
		}
		byteMap[s[j]] = j
	}

	return maxLen
}
