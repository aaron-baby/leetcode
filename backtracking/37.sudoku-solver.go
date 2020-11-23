package backtracking

func solveSudoku(board [][]byte) {
	helper(board, 0, 0)
}

func helper(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	// move to next row
	if j == 9 {
		return helper(board, i+1, 0)
	}
	// move to next cell
	if board[i][j] != '.' {
		return helper(board, i, j+1)
	}
	var c byte
	for c = '1'; c <= '9'; c++ {
		if !isValid(board, i, j, c) {
			continue
		}
		board[i][j] = c
		if helper(board, i, j+1) {
			return true
		}
		// back track
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, i, j int, c byte) bool {
	// check column
	for x := 0; x < 9; x++ {
		if board[x][j] == c {
			return false
		}
	}
	// check row
	for y := 0; y < 9; y++ {
		if board[i][y] == c {
			return false
		}
	}
	// check subbox
	row, col := i-i%3, j-j%3
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if board[row+x][col+y] == c {
				return false
			}
		}
	}
	return true
}
