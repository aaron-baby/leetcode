package array

func spiralOrder(matrix [][]int) []int {
	var res = []int{}
	if len(matrix) == 0 {
		return res
	}
	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		// top c1->c2
		for c := c1; c <= c2; c++ {
			res = append(res, matrix[r1][c])
		}
		// right r1+1->r2
		for r := r1 + 1; r <= r2; r++ {
			res = append(res, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			// c2-1 -> c1+1
			for c := c2 - 1; c > c1; c-- {
				res = append(res, matrix[r2][c])
			}
			// r2 -> r1+1
			for r := r2; r > r1; r-- {
				res = append(res, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return res
}
