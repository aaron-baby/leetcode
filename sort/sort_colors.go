package sort

// 75. Sort Colors
// https://leetcode.com/problems/sort-colors/
func sortColors(nums []int) {
	if len(nums) <= 1 {
		return
	}
	var start, end = 0, len(nums) - 1
	for i := 0; i <= end; i++ {
		// swap current to end
		if nums[i] == 2 {
			nums[end], nums[i] = nums[i], nums[end]
			end--
			i--
		} else if nums[i] == 0 {
			nums[start], nums[i] = nums[i], nums[start]
			start++
		}
	}
}
