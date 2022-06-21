package math

import "strconv"

func maxPoints(points [][]int) int {
	var x = map[int]int{}
	var res int
	if len(points) <= 2 {
		res = len(points)
	}
	for i, p := range points {
		_, found := x[p[0]]
		if !found {
			x[p[0]] = 0
		}
		x[p[0]] += 1
		if x[p[0]] > res {
			res = x[p[0]]
		}
		var slope = map[string]int{}
		// linear
		for j := 0; j < len(points); j++ {
			// horizon
			if points[j][0] == points[i][0] || i == j {
				continue
			}
			dx, dy := points[j][0]-points[i][0], points[j][1]-points[i][1]
			g := gcd(dx, dy)
			s := strconv.Itoa(dy/g) + "_" + strconv.Itoa(dx/g)
			_, found = slope[s]
			if !found {
				slope[s] = 1
			}
			slope[s] += 1
			if slope[s] > res {
				res = slope[s]
			}
		}
	}
	return res
}
func gcd(a int, b int) int {
	for b != 0 {
		// remainder
		a = a % b
		a, b = b, a
	}
	return a
}
