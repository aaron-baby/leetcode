package two_pointers

func trap(height []int) int {
	var left, right = 0, len(height)-1
	ans :=0
	leftMax, rightMax :=0,0
	for left<right {
		// 右高左低
		if height[left]<height[right] {
			if height[left]>= leftMax {
				leftMax = height[left]
			} else {
				ans += leftMax - height[left]
			}
			left++
		} else {
			if height[right]>= rightMax {
				rightMax = height[right]
			} else {
				ans += rightMax - height[right]
			}
			right--
		}
	}
	return ans
}
