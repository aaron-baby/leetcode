package dynamic_programming

import "gitlab.com/aaw/leetcode/lib"

func coinChange(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i >= coins[j] {
				dp[i] = lib.Min(dp[i-coins[j]]+1, dp[i])
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
