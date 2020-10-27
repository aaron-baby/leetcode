package backtracking

func combine(n int, k int) [][]int {
	if k > n || k <= 0 || n <= 0 {
		return [][]int{}
	}
	c, res := []int{}, [][]int{}
	generateCombination(n, k, 1, c, &res)
	return res
}

func generateCombination(n, k, start int, c []int, res *[][]int) {
	if len(c) == k {
		b := make([]int, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	// k-len(c) is the number count that we needed
	// (n - i +1) length from opposite side should greater than or equal requirement
	for i := start; i <= n-(k-len(c))+1; i++ {
		c = append(c, i)
		generateCombination(n, k, i+1, c, res)
		c = c[:len(c)-1]
	}
}
