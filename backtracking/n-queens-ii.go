package backtracking

func totalNQueens(n int) int {
	rows := make([]int, n)
	dales := make([]int, 4*n-1)
	hills := make([]int, 2*n-1)
	return backtrack52(0, 0, n, rows, hills, dales)
}

func notUnderAttack(row, col, n int, rows, hills, dales []int) bool {
	res := rows[col] + dales[row-col+2*n] + hills[row+col]
	if res > 0 {
		return false
	}
	return true
}

func backtrack52(row, count, n int, rows, hills, dales []int) int {
	for col := 0; col < n; col++ {
		if notUnderAttack(row, col, n, rows, hills, dales) {
			// place queen, column col is occupy
			rows[col] = 1
			dales[row-col+2*n] = 1
			hills[row+col] = 1

			if row+1 == n {
				count++
			} else {
				count = backtrack52(row+1, count, n, rows, hills, dales)
			}

			// remove queen
			rows[col] = 0
			dales[row-col+2*n] = 0
			hills[row+col] = 0
		}
	}
	return count
}
