package greedy_algorithm

import (
	"gitlab.com/aaw/leetcode/lib"
)

func candy(ratings []int) int {
	var nums = make([]int, len(ratings))
	for i := range nums {
		nums[i] = 1
	}
	for i := 0; i < len(ratings)-1; i++ {
		if ratings[i+1] > ratings[i] {
			nums[i+1] = nums[i] + 1
		}
	}
	for i := len(ratings) - 1; i > 0; i-- {
		if ratings[i-1] > ratings[i] {
			nums[i-1] = lib.Max(nums[i-1], nums[i]+1)
		}
	}
	var res int
	for _, v := range nums {
		res += v
	}
	return res
}
