package backtracking

func exist(board [][]byte, word string) bool {
	if len(board) == 0 {
		return false
	}
	visited := make(map[[2]int]struct{})
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if visitCell(board, word, i, j, 0, visited) {
				return true
			}
		}
	}
	return false
}

func visitCell(board [][]byte, word string, i, j, idx int, visited map[[2]int]struct{}) bool {
	if idx == len(word) {
		return true
	}
	m, n := len(board), len(board[0])
	_, ok := visited[[2]int{i, j}]
	if ok {
		return false
	}
	if i < 0 || j < 0 || i > m-1 || j > n-1 || board[i][j] != word[idx] {
		return false
	}
	visited[[2]int{i, j}] = struct{}{}
	res := visitCell(board, word, i+1, j, idx+1, visited) ||
		visitCell(board, word, i-1, j, idx+1, visited) ||
		visitCell(board, word, i, j-1, idx+1, visited) ||
		visitCell(board, word, i, j+1, idx+1, visited)
	delete(visited, [2]int{i, j})
	return res
}
