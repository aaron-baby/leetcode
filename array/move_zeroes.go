package array

func moveZeroes(nums []int) {
	// non zero pointer
	var cnt = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[cnt] = nums[i]
			cnt++
		}
	}
	for cnt < len(nums) {
		nums[cnt] = 0
		cnt++
	}
}
