package dynamic_programming

import "gitlab.com/aaw/leetcode/lib"

func rob(nums []int) int {
	// rob(i) = Math.max( rob(i - 2) + currentHouseValue, rob(i - 1) )
	// A robber has 2 options: a) rob current house i; b) don't rob current house.
	// option "b") loot from the previous house robbery and any loot captured before that

	if len(nums) == 0 {
		return 0
	}
	// bottom-up
	dp := make([]int, len(nums)+1)
	dp[0], dp[1] = 0, nums[0]
	for i := 2; i <= len(nums); i++ {
		dp[i] = lib.Max(dp[i-2]+nums[i-1], dp[i-1])
	}
	return dp[len(nums)]
}
