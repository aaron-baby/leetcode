package main

// 33. Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/

func search(nums []int, target int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		m := left + (right-left)/2
		if nums[m] == target {
			return m
		}
		// right half in ascending order
		if nums[m] < nums[right] {
			// target in left half
			if target > nums[m] && target <= nums[right] {
				// move left cursor
				left = m + 1
			} else {
				right = m - 1
			}
		} else {
			if target >= nums[left] && target < nums[m] {
				right = m - 1
			} else {
				left = m + 1
			}
		}
	}

	return -1
}
