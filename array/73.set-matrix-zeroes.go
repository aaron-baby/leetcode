package array

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	var firstRowZero, firstColZero = false, false
	r, c := len(matrix), len(matrix[0])
	for i := 0; i < r; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
		}
	}
	for j := 0; j < c; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
		}
	}
	// use first row and column to set zero flag
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 1; i < r; i++ {
		for j := 1; j < c; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}
	if firstRowZero {
		for j := 0; j < c; j++ {
			matrix[0][j] = 0
		}
	}
	if firstColZero {
		for i := 0; i < r; i++ {
			matrix[i][0] = 0
		}
	}
}
