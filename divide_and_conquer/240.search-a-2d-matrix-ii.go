package divide_and_conquer

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	// start from top right corner
	for i, j := 0, n-1; i < m && j >= 0; {
		if target == matrix[i][j] {
			return true
		}
		// reduce column
		if target < matrix[i][j] {
			j--
		} else {
			// look on this column
			i++
		}
	}
	return false
}
