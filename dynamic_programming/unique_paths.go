package dynamic_programming

// 62. Unique Paths
// https://leetcode.com/problems/unique-paths/
var mem = make(map[[2]int]int)
func uniquePaths(m int, n int) int {
	// only one path to this point
	if m == 1 || n == 1{
		return 1
	}
	v, ok := mem[[2]int{m,n}]
	if !ok {
		v = uniquePaths(m-1, n)+uniquePaths(m,n-1)
		mem[[2]int{m,n}] = v
	}

	return v
}
