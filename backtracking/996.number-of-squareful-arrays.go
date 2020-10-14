package backtracking

import (
	"math"
	"sort"
)

func numSquarefulPerms(A []int) int {
	if len(A) == 0 {
		return 0
	}
	used, p, res := make([]bool, len(A)), []int{}, [][]int{}
	sort.Ints(A)
	generatePermutation(A, 0, p, &res, &used)
	return len(res)
}

func generatePermutation(nums []int, index int, p []int, res *[][]int, used *[]bool) {
	if index == len(nums) {
		temp := make([]int, len(p))
		copy(temp, p)
		*res = append(*res, temp)
	}
	for i := 0; i < len(nums); i++ {
		if (*used)[i] {
			continue
		}
		// prev number not qualify
		if i > 0 && nums[i] == nums[i-1] && !(*used)[i-1] {
			continue
		}
		// with prev number sum is not square
		if len(p) > 0 && !checkSquare(nums[i]+p[len(p)-1]) {
			continue
		}
		(*used)[i] = true
		p = append(p, nums[i])
		generatePermutation(nums, index+1, p, res, used)
		p = p[:len(p)-1]
		(*used)[i] = false
	}
}
func checkSquare(num int) bool {
	tmp := math.Sqrt(float64(num))
	if int(tmp)*int(tmp) == num {
		return true
	}
	return false
}
