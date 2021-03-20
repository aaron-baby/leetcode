package greedy_algorithm

func isSubsequence(s string, t string) bool {
	// initial pointer
	i, j := 0, 0
	for i < len(s) && j < len(t) {
		if t[j] == s[i] {
			i++
		}
		j++
	}
	return i == len(s)
}
