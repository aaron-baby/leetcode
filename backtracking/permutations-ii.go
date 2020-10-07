package backtracking

// 47. Permutations II
// https://leetcode.com/problems/permutations-ii/
func permuteUnique(nums []int) [][]int {
	counter := make(map[int]int)
	for _, v := range nums {
		counter[v] += 1
	}
	results := [][]int{}
	back_track([]int{}, len(nums), &counter, &results)
	return results
}

func back_track(comb []int, n int, counter *map[int]int, results *[][]int) {
	if len(comb) == n {
		row := make([]int, n)
		copy(row, comb)
		*results = append(*results, row)
		return
	}
	c := *counter
	for num, count := range c {
		if count == 0 {
			continue
		}
		comb = append(comb, num)
		c[num] -= 1

		back_track(comb, n, counter, results)

		comb = comb[:len(comb)-1]
		c[num] += 1
	}
}
