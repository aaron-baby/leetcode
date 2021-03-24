package array

import "gitlab.com/aaw/leetcode/lib"

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var cnt int
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if lib.Abs(arr[i]-arr[j]) <= a && lib.Abs(arr[j]-arr[k]) <= b && lib.Abs(arr[i]-arr[k]) <= c {
					cnt++
				}
			}
		}
	}
	return cnt
}
