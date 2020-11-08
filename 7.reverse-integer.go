package main

func reverse(x int) int {
	var r int
	for x != 0 {
		r = r*10 + x%10
		// quotient
		x /= 10
	}
	if r < -1<<31 || r > 1<<31-1 {
		return 0
	}
	return r
}
