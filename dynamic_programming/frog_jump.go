package dynamic_programming

import (
	"fmt"
	"sort"
)

// 403. Frog Jump
// https://leetcode.com/problems/frog-jump/
func canCross(stones []int) bool {
	dp := make(map[int]map[int]int, len(stones))
	for _, v := range stones {
		dp[v] = make(map[int]int)
	}
	dp[0][0] = 0
	for i := 0; i < len(stones); i++ {
		for _, k := range dp[stones[i]] {
			for step := k - 1; step <= k+1; step++ {
				if step > 0 {
					if _, ok := dp[stones[i]+step]; ok {
						dp[stones[i]+step][i] = step
					}
				}
			}
		}
	}
	// printDP(dp)
	return len(dp[stones[len(stones)-1]]) != 0
}

func printDP(dp map[int]map[int]int) {
	keys := []int{}
	for k := range dp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Println(k, dp[k])
	}

}
