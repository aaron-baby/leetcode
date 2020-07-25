package dynamic_programming

// 416. Partition Equal Subset Sum
// https://leetcode.com/problems/partition-equal-subset-sum/
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 != 0 {
		return false
	}
	target, n := sum/2, len(nums)

	// dp(i, j) be true if a subset of { x1, ..., xj } sums to i and false otherwise.
	// Then dp(half_sum, n) is true if and only if there is a subset of S that sums to half_sum.
	dp := make([][]bool, target+1)
	for r := range dp {
		dp[r] = make([]bool, n+1)
		if r > 0 {
			dp[r][0] = false
		}
	}
	for i := 0; i < n+1; i++ {
		dp[0][i] = true
	}
	for i := 1; i <= target; i++ {
		for j := 1; j <= n; j++ {
			if i-nums[j-1] >= 0 {
				dp[i][j] = dp[i][j-1] || dp[i-nums[j-1]][j-1]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}
	return dp[target][n]
}
