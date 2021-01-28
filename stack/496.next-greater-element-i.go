package stack

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var s []int
	var m = make(map[int]int)
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums2[i] {
			s = s[:len(s)-1] // pop stack elements small than current
		}
		v := -1
		if len(s) > 0 {
			v = s[len(s)-1]
		}
		m[nums2[i]] = v

		s = append(s, nums2[i])
	}
	var rst []int
	for _, v := range nums1 {
		rst = append(rst, m[v])
	}
	return rst
}
