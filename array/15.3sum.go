package array

import "sort"

func threeSum(nums []int) [][]int {
	var res [][]int
	// Run a nested loop with two loops, outer loop from 0 to n-2 and the inner loop from i+1 to n-1
	picked := make(map[[3]int]struct{},len(nums))

	for i := 0; i < len(nums)-1; i++ {
		m := make(map[int]int, len(nums))
		for j := i + 1; j < len(nums); j++ {
			diff := -(nums[i]+nums[j])
			if _, ok := m[diff]; ok {
				var a [3]int
				s := []int{nums[i], nums[j], diff}
				sort.Ints(s)
				copy(a[:], s)
				if _, ok := picked[a]; ok{
					continue
				}
				picked[a] = struct{}{}
				res = append(res, s)
			} else {
				m[nums[j]] = 1
			}
		}
	}
	return res
}
