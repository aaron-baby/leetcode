package two_pointers

import "gitlab.com/aaw/leetcode/lib"

func maxDistToClosest(seats []int) int {
	left := -1
	var right, maxDist int
	ans := 0
	for i := 0; i < len(seats); i++ {
		if seats[i] == 1 {
			left = i
			continue
		}
		for right < len(seats) && seats[right] == 0 || right < i {
			right++
		}
		// right distance
		leftD := i - left
		if left == -1 {
			leftD = len(seats)
		}
		rightD := right - i
		if right == len(seats) {
			rightD = len(seats)
		}

		maxDist = lib.Min(leftD, rightD)
		if maxDist > ans {
			ans = maxDist
		}
	}
	return ans
}
