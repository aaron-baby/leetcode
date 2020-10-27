package two_pointers

import "gitlab.com/aaw/leetcode/lib"

func maxArea(height []int) int {
	maxarea, l, r := 0, 0, len(height)-1
	for l < r {
		maxarea = lib.Max(maxarea, lib.Min(height[l], height[r])*(r-l))
		// If we try to move the pointer at the longer line inwards,
		// we won't gain any increase in area, since it is limited by the shorter line.
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return maxarea
}
