package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/lib"
	"math"
)

func maxProfitCooldown(prices []int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	// base case
	dp[0][0] = 0
	dp[0][1] = math.MinInt32
	pre0 := 0
	for i := 1; i <= n; i++ {
		dp[i][0] = lib.Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// After you sell your stock, you cannot buy stock on next day.
		// T[i][k][0] and T[i][k][1], where the former denotes
		// the maximum profit at the end of the i-th day with at most k transactions and with 0 stock in our hand AFTER taking the action
		if i >= 2 {
			pre0 = dp[i-2][0]
		}
		dp[i][1] = lib.Max(dp[i-1][1], pre0-prices[i-1])
	}
	return dp[n][0]
}
