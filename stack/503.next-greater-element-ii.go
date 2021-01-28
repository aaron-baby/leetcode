package stack

func nextGreaterElements(nums []int) []int {
	var s []int
	n := len(nums)
	var rst = make([]int, n)
	for i := (n - 1) * 2; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1] // pop stack elements small than current
		}
		v := -1
		if len(s) > 0 {
			v = s[len(s)-1]
		}
		rst[i%n] = v
		s = append(s, nums[i%n])
	}
	return rst
}
