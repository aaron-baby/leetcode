package dynamic_programming

func isMatch(s string, p string) bool {
	// dp[i][j] stands for whether s[0:i-1] matches p[0:j-1]
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	// dp[i][j] = dp[i][j - 2], if p[j - 1] == '*' and the pattern repeats for 0 time
	// dp[i][j] = dp[i - 1][j] && (s[i - 1] == p[j - 2] || p[j - 2] == '.'), if p[j - 1] == '*' and the pattern repeats for at least 1 time
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if j > 1 && p[j-1] == '*' {
				dp[i][j] = dp[i][j-2] || (i > 0 && dp[i-1][j] && (s[i-1] == p[j-2] || p[j-2] == '.'))
			} else {
				// dp[i][j] = dp[i - 1][j - 1], if p[j - 1] != '*' && (s[i - 1] == p[j - 1] || p[j - 1] == '.')
				dp[i][j] = i > 0 && dp[i-1][j-1] && (s[i-1] == p[j-1] || p[j-1] == '.')
			}
		}
	}
	return dp[m][n]
}
