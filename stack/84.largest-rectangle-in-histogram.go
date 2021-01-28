package stack

import (
	"gitlab.com/aaw/leetcode/lib"
)

func largestRectangleArea(heights []int) int {
	// index of the first bar from left lower than current
	left := make([]int, len(heights))
	left[0] = -1
	for i := 1; i <= len(heights)-1; i++ {
		pre := i - 1
		// 前一个元素比当前元素大
		for pre >= 0 && heights[pre] >= heights[i] {
			pre = left[pre]
		}
		left[i] = pre
	}
	right := make([]int, len(heights))
	right[len(heights)-1] = len(heights)
	for i := len(heights) - 1; i >= 0; i-- {
		after := i + 1
		// 后一个元素比当前元素大
		for after <= len(heights)-1 && heights[after] >= heights[i] {
			after = right[after]
		}
		right[i] = after
	}
	var maxArea int
	for i := 0; i < len(heights); i++ {
		maxArea = lib.Max(maxArea, heights[i]*(right[i]-left[i]-1))
	}
	return maxArea
}
