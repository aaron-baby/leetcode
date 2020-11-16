package main

func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	temp := myPow(x, n/2)
	if n%2 == 0 {
		return temp * temp
	}
	// odd case
	if n > 0 {
		return x * temp * temp
	} else {
		return (temp * temp) / x
	}
}
