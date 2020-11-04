package greedy_algorithm

import "gitlab.com/aaw/leetcode/lib"

func jump(nums []int) int {
	var steps int
	last, curr := 0, 0
	for i := 0; i < len(nums); i++ {
		// position beyond last reach
		if i > last {
			last = curr
			steps++
		}
		// furthest distance can be reach
		curr = lib.Max(curr, i+nums[i])
	}
	return steps
}
