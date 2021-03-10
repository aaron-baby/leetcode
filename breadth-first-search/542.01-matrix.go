package breadth_first_search

import "gitlab.com/aaw/leetcode/lib"

func updateMatrix(matrix [][]int) [][]int {
	var rows = len(matrix)
	if rows == 0 {
		return matrix
	}
	var cols = len(matrix[0])
	var dist = make([][]int, rows)
	max := lib.Max(rows, cols)
	for i := range dist {
		dist[i] = make([]int, cols)
		for j := 0; j < cols; j++ {
			dist[i][j] = max
		}
	}
	// For our BFS routine, we keep a queue, q to maintain the queue of cells to be examined next.
	var q [][2]int
	// We start by adding all the cells with 0s to q.
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				dist[i][j] = 0
				q = append(q, [2]int{i, j})
			}
		}
	}
	var curr [2]int
	// Pop the cell from queue, and examine its neighbours.
	for len(q) > 0 {
		curr, q = q[0], q[1:]
		dir := [4][2]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		for _, o := range dir {
			nr, nc := curr[0]+o[0], curr[1]+o[1]
			// check boundaries
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
				// remaining nodes has value 1
				if dist[curr[0]][curr[1]]+1 < dist[nr][nc] {
					dist[nr][nc] = dist[curr[0]][curr[1]] + 1
					q = append(q, [2]int{nr, nc})
				}
			}
		}
	}
	return dist
}
