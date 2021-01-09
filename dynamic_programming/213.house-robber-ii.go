package dynamic_programming

import "gitlab.com/aaw/leetcode/lib"

func robii(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// For example, 1 -> 2 -> 3 -> 1 becomes 2 -> 3 if 1 is not robbed.
	// Since every house is either robbed or not robbed,
	// let's say house 0 is not robbed, it's can be simply to solve subproblem 1..n-1
	// house 0 is robbed, then previous consecutive house n-1 must not be robbed, we can use it to break the circle
	// that is say to solve subproblem 0..n-2
	return lib.Max(rob(nums[1:]), rob(nums[:len(nums)-1]))
}
