package dynamic_programming

import (
	"gitlab.com/aaw/leetcode/lib"
	"math"
)

func numSquares(n int) int {
	d := n + 1
	if d < 4 {
		d = 4
	}
	dp := make([]int, d)
	// 	   1 	  1+1
	dp[0], dp[1], dp[2] = 0, 1, 2
	// 1+1+1
	dp[3] = 3
	for i := 4; i <= n; i++ {
		// worst scenario
		dp[i] = i
		c := int(math.Ceil(math.Sqrt(float64(n))))
		for x := 1; x <= c; x++ {
			temp := x * x
			if temp > i {
				break
			}
			dp[i] = lib.Min(dp[i], 1+dp[i-temp])
		}
	}

	return dp[n]
}
