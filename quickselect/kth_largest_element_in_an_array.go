package quickselect

// 215. Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	// pick a pivot
	pivot := nums[0]
	// split into a pile A1 of small elements and A2 of big elements
	var a1, a2 []int
	for i := 0; i < len(nums); i++ {
		if nums[i] < pivot {
			a1 = append(a1, nums[i])
		} else if nums[i] > pivot {
			a2 = append(a2, nums[i])
		}
	}

	if k <= len(a2) {
		return findKthLargest(a2, k)
	} else if len(nums)-k < len(a1) {
		// k in small pile, have excluded total - small pile size
		return findKthLargest(a1, k-(len(nums)-len(a1)))
	}

	return pivot
}
