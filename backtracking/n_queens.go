package backtracking

var leftDiagonal, rightDiagonal, cl []int
// 51. N-Queens
// https://leetcode.com/problems/n-queens/
func solveNQueens(n int) [][]string {
	slt := [][]string{}
	leftDiagonal, rightDiagonal, cl = make([]int, 2*n-1), make([]int, 2*n-1), make([]int, n)
	board := make([][]int, n)
	// column left to right
	for j := 0; j < n; j++ {
		board[j] = make([]int, n)
	}
	solveNQueenRecursive(board, 0, &slt)

	return slt
}

func printBoard(board [][]int) []string {
	var r []string
	for i := 0; i < len(board); i++ {
		t := ""
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 1 {
				t += "Q"
			} else {
				t += "."
			}
		}
		r = append(r, t)
	}

	return r
}

func solveNQueenRecursive(board [][]int, col int, slt *[][]string) bool {
	n := len(board)
	if col >= n {
		return true
	}
	var res bool
	for i := 0; i < n; i++ {
		if leftDiagonal[i-col+n-1] != 1 && rightDiagonal[i+col] != 1 && cl[i] != 1 {
			board[i][col] = 1
			leftDiagonal[i-col+n-1], rightDiagonal[i+col] = 1, 1
			cl[i] = 1

			res = solveNQueenRecursive(board, col+1, slt)
			if res {
				*slt = append(*slt, printBoard(board))
			}

			// back track
			board[i][col] = 0
			leftDiagonal[i-col+n-1], rightDiagonal[i+col] = 0, 0
			cl[i] = 0
		}
	}

	return false
}
