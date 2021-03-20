package union_find

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	// m for row size, n for column size
	var m, n = len(board), len(board[0])
	for y := 0; y <= n-1; y++ {
		dfs130(board, 0, y, m, n)
		dfs130(board, m-1, y, m, n)
	}
	for x := 0; x <= m-1; x++ {
		dfs130(board, x, 0, m, n)
		dfs130(board, x, n-1, m, n)
	}
	for x := 0; x < m; x++ {
		for y := 0; y < n; y++ {
			if board[x][y] == 'A' {
				board[x][y] = 'O'
			} else {
				board[x][y] = 'X'
			}
		}
	}
}

// visit boarder 'O's
func dfs130(board [][]byte, x, y, m, n int) {
	// out of boundary
	if x >= m || x < 0 || y >= n || y < 0 || board[x][y] != 'O' {
		return
	}
	board[x][y] = 'A'
	dfs130(board, x+1, y, m, n)
	dfs130(board, x-1, y, m, n)
	dfs130(board, x, y-1, m, n)
	dfs130(board, x, y+1, m, n)
}
