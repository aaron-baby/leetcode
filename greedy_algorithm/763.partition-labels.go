package greedy_algorithm

func partitionLabels(S string) []int {
	last := make(map[uint8]int, 26)
	for i := 0; i < len(S); i++ {
		last[S[i]-'a'] = i
	}
	anchor, j := 0, 0
	var ans []int
	for i := 0; i < len(S); i++ {
		if last[S[i]-'a'] > j {
			j = last[S[i]-'a']
		}
		if i == j {
			ans = append(ans, i-anchor+1)
			anchor = i + 1
		}
	}
	return ans
}
