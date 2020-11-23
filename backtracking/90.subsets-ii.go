package backtracking

import (
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	c, res := []int{}, [][]int{}
	sort.Ints(nums)
	for k := 0; k <= len(nums); k++ {
		generateSubset90(nums, k, 0, c, &res)
	}
	return res
}

func generateSubset90(nums []int, k int, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// k-len(c) is the number count that we needed
	// (n - i +1) length from opposite side should greater than or equal requirement
	// index start from 0
	for i := start; i <= len(nums)-1-(k-len(c))+1; i++ {
		if i > start && nums[i] == nums[i-1] {
			continue
		}
		c = append(c, nums[i])
		generateSubset90(nums, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}
