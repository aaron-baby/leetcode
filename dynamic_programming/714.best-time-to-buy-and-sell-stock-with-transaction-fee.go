package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/lib"
	"math"
)

func maxProfitFee(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = math.MinInt32
	for i := 1; i <= n; i++ {
		dp[i][0] = lib.Max(dp[i-1][0], dp[i-1][1]+prices[i-1])
		// subtract transaction fee on buying
		dp[i][1] = lib.Max(dp[i-1][1], dp[i-1][0]-prices[i-1]-fee)
	}
	return dp[n][0]
}
