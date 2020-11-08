package dynamic_programming

func numDecodings(s string) int {
	// dp stands for number of solutions that decode a string's length equal to n
	dp := make(map[int]int, len(s)+1)
	// an empty string decode to empty
	dp[0] = 1
	if len(s) > 0 && s[0] == '0' {
		return 0
	}
	// without leading zero number can be decode
	dp[1] = 1
	for i := 2; i <= len(s); i++ {
		// the last digit can decode to letter
		if s[i-1] > '0' {
			dp[i] += dp[i-1]
		}
		// last two digits can alternatively decode to a valid letter
		if s[i-2] == '1' || (s[i-2] == '2' && s[i-1] < '7') {
			dp[i] += dp[i-2]
		}
	}
	return dp[len(s)]
}
