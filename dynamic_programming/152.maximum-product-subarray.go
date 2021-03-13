package dynamic_programming

import "gitlab.com/aaw/leetcode/lib"

func maxProduct(nums []int) int {
	// consider negative scenario
	maximum, minimum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maximum, minimum = minimum, maximum
		}
		maximum = lib.Max(maximum*nums[i], nums[i])
		minimum = lib.Min(minimum*nums[i], nums[i])
		res = lib.Max(res, maximum)
	}
	return res
}
