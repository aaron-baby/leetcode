package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/lib"
)

func minimumTotal(triangle [][]int) int {
	var r = len(triangle)
	dp := make([][]int, r)
	for i := 0; i < r; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	// initial value of bottom row
	for i, v := range triangle[r-1] {
		dp[r-1][i] = v
	}
	for i := r - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = triangle[i][j] + lib.Min(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}
