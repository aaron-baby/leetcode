package main

// 41. First Missing Positive
// Your algorithm should run in O(n) time and uses constant extra space.
func firstMissingPositive(nums []int) int {
	n := len(nums)

	// sequence positive number should be range from 1 to n, otherwise it out of range
	for i := 0; i < n; i++ {
		// e.g. number 5 should be at index (5-1), cause index are start from 0
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			i--
		}
	}
	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
