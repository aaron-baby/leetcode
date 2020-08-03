package array

// 238. Product of Array Except Self
// https://leetcode.com/problems/product-of-array-except-self/
func productExceptSelf(nums []int) []int {
	l := len(nums)
	left, right, ans := make([]int, l), make([]int, l), make([]int, l)
	left[0] = 1
	for i := 1; i < l; i++ {
		left[i] = nums[i-1] * left[i-1]
	}
	right[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}
	for i := 0; i < l; i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
