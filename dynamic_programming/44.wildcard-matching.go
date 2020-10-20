package dynamic_programming

// We can use Dynamic Programming to solve this problem –
// Let dp[i][j] is true if first i characters in given string matches the first j characters of pattern.
func isMatch44(s string, p string) bool {
	// DP Initialization:
	// both text and pattern are null
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true

	// pattern is null
	for i := 1; i < m+1; i++ {
		dp[i][0] = false
	}

	// text is null
	for j := 1; j < n+1; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if p[j-1] == '*' {
				// a) * indicates an empty sequence or b) * match ith
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			} else if p[j-1] == '?' || s[i-1] == p[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}
