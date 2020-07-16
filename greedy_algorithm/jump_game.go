package greedy_algorithm

// 55. Jump Game
func canJump(nums []int) bool {
	var lastPos = len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		// current index plus max steps can reach to last position, initially it's last index
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
