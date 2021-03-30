package dynamic_programming

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	len := 2
	res := 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			len++
		} else {
			if len > 2 {
				res += (len - 1) * (len - 2) / 2
			}
			len = 2
		}
	}
	if len > 2 {
		res += (len - 1) * (len - 2) / 2
	}
	return res
}
