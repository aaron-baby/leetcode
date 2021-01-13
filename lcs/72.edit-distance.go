package lcs

import "gitlab.com/aaw/leetcode/lib"

func minDistance(word1 string, word2 string) int {
	// To apply DP, we define the state dp[i][j] to be the minimum number of operations to convert word1[0..i) to word2[0..j).

	// Suppose we have already known how to convert word1[0..i - 1) to word2[0..j - 1) (dp[i - 1][j - 1]),
	// if word1[i - 1] == word2[j - 1], then no more operation is needed and dp[i][j] = dp[i - 1][j - 1].

	// If word1[i - 1] != word2[j - 1], we need to consider three cases.
	// 		Replace word1[i - 1] by word2[j - 1] (dp[i][j] = dp[i - 1][j - 1] + 1);
	// If word1[0..i - 1) = word2[0..j) then delete word1[i - 1]
	//	dp[i][j] = dp[i-1][j]+1

	// If word1[0..i) + word2[j - 1] = word2[0..j) then insert word2[j - 1] to word1[0..i)
	// dp[i][j] = dp[i][j-1]+1
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
		dp[i][0] = i
	}
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
				continue
			}
			dp[i][j] = lib.Min(lib.Min(dp[i-1][j-1]+1, dp[i-1][j]+1), dp[i][j-1]+1)
		}
	}
	return dp[len(word1)][len(word2)]
}
