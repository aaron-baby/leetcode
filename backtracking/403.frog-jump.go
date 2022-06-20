package backtracking

func canCross(stones []int) bool {
	var step = map[int]int{}
	// dp[i][j] store does it possible jump from position i to j
	dp := map[[2]int]bool{}
	return canJump(0, 0, len(stones), step, stones, dp)
}
func canJump(p int, c int, l int, lastJump map[int]int, stones []int, dp map[[2]int]bool) bool {
	// p previous position
	// c current position
	// l total length
	// last_jump last jump k units

	// landing on the last stone
	if c == l-1 {
		return true
	}
	if reachable, ok := dp[[2]int{p, c}]; ok {
		//log.Printf("hit jump from %v to %v", p, c)
		return reachable
	}

	// Initially, the frog is on the first stone and assumes the first jump must be 1 unit.
	var lowBound, highBound int
	if c == 0 {
		lowBound, highBound = 1, 1
	} else {
		k := lastJump[p]
		lowBound, highBound = k-1, k+1
	}
	// farthest stone can be reach
	for n := c + 1; n < l && stones[n]-stones[c] <= highBound; n++ {
		// unable move to next stone
		if stones[n]-stones[c] < lowBound {
			continue
		}
		//fmt.Printf("n: %v, c: %v, diff units: %v\n", n, c, stones[n]-stones[c])
		lastJump[c] = stones[n] - stones[c]
		// move to next stone
		branchRes := canJump(c, n, l, lastJump, stones, dp)
		if branchRes {
			dp[[2]int{c, n}] = true
			return branchRes
		} else {
			dp[[2]int{c, n}] = false
		}
		//backtrack
		delete(lastJump, c)
	}
	return false
}
