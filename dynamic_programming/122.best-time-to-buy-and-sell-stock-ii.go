package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/lib"
	"math"
)

func maxProfit(prices []int) int {
	// T[i][k] be the maximum profit that could be gained at the end of the i-th day with at most k transactions.
	// The number of stocks held in our hand is the hidden factor mentioned above that will affect the action on the i-th day and thus affect the maximum profit.

	// Recurrence relations:
	// T[i][k][0] = max(T[i-1][k][0], T[i-1][k][1] + prices[i])
	// 昨天没持有，今天没有买卖 或者 昨天持有今天卖了
	// T[i][k][1] = max(T[i-1][k][1], T[i-1][k-1][0] - prices[i])
	// 昨天持有，今天无买卖 或者 昨天没持有今天买入

	// Base cases:第一天之前的值
	// T[-1][k][0] = 0, T[-1][k][1] = -Infinity

	// k equal to 0 means no transaction
	//	T[i][0][0] = 0, T[i][0][1] = -Infinity

	// k no limitation
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		dp[i][0] = lib.Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		dp[i][1] = lib.Max(dp[i-1][1], dp[i-1][0]-prices[i-1])
	}
	return dp[n][0]
}
