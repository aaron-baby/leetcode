package array

// 34. Find First and Last Position of Element in Sorted Array
// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	start := firstGreaterEqual(nums, target)
	if len(nums) == 0 || start == len(nums) || nums[start] != target {
		return []int{-1, -1}
	}
	end := firstGreaterEqual(nums, target+1) - 1

	return []int{start, end}
}
func firstGreaterEqual(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
