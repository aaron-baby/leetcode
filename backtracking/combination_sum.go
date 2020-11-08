package backtracking

import "fmt"

// 39. Combination Sum
// https://leetcode.com/problems/combination-sum/
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target == 0 {
		fmt.Println(c)
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := index; i < len(nums); i++ {
		fmt.Println("- - ", index)
		fmt.Printf("i:%d, %d  target:%d\n", i, nums[i], target)
		if nums[i] > target {
			continue
		}
		fmt.Println("pick", nums[i])
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res)
		fmt.Println("drop", c[len(c)-1])
		c = c[:len(c)-1]
	}
}
