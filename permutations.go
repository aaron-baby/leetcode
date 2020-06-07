package main

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	h, tail := nums[0], nums[1:]
	res := make([][]int, 0)

	//fmt.Printf("%d,%+v\n===\n", h, tail)
	for _, v := range permute(tail) {
		for j := 0; j <= len(v); j++ {
			//fmt.Println(h, j, v[:j], v[j:])
			s := make([]int, len(v[:j]))
			copy(s, v[:j])
			s = append(append(s, h), v[j:]...)
			//fmt.Println(s,"\n---")
			res = append(res, s)
		}
	}

	return res
}
