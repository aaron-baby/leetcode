package hash_table

func isValidSudoku(board [][]byte) bool {
	row, column, subboxes := make([]map[byte]int, 9), make([]map[byte]int, 9), make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		row[i] = make(map[byte]int)
		column[i] = make(map[byte]int)
		subboxes[i] = make(map[byte]int)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if '.' == board[i][j] {
				continue
			}
			if _, ok := row[i][board[i][j]]; ok {
				return false
			} else {
				row[i][board[i][j]] = 1
			}

			if _, ok := column[j][board[i][j]]; ok {
				return false
			} else {
				column[j][board[i][j]] = 1
			}

			// 0|1|2
			// 3|4|5
			// 6|7|8
			boxIndex := (i/3)*3 + j/3
			if _, ok := subboxes[boxIndex][board[i][j]]; ok {
				return false
			} else {
				subboxes[boxIndex][board[i][j]] = 1
			}
		}
	}
	return true
}
