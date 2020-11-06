package backtracking

import (
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum40(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum40(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		// do not select same number in this loop, only recursive
		if i > index && nums[i] == nums[i-1] {
			continue
		}
		if nums[i] > target || target == 0 {
			continue
		}
		c = append(c, nums[i])
		findCombinationSum40(nums, target-nums[i], i+1, c, res)
		c = c[:len(c)-1]
	}
}
