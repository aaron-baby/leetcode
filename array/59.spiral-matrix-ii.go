package array

func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	u, d := 0, n-1
	l, r := 0, n-1
	for i := 1; i <= n*n; {
		for col := l; col <= r; col++ {
			res[u][col] = i
			i++
		}
		u++
		if u > d {
			break
		}
		for row := u; row <= d; row++ {
			res[row][r] = i
			i++
		}
		r--
		if r < l {
			break
		}
		for col := r; col >= l; col-- {
			res[d][col] = i
			i++
		}
		d--
		if d < u {
			break
		}
		for row := d; row >= u; row-- {
			res[row][l] = i
			i++
		}
		l++
		if l > r {
			break
		}
	}
	return res
}
