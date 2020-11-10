package dynamic_programming

func numTrees(n int) int {
	// to calculate unique bst, we start from simplest case
	// dp[n] stands for n node scenario, dp[0] is an empty tree, equal to 1
	// dp[1] can only have 1 node as root
	// dp[2] pick 1 as root, left tree is empty, right tree have one node
	// 	   pick 2 as root, left tree has one node, right empty
	//     left trees combination * right trees combination
	//     = dp[0] * dp[1] + dp[1] * dp[0]
	// dp[3] can each pick 1,2,3 as root
	//     = dp[0] * dp[2]
	//     + dp[1] * dp[1]
	//     + dp[2] * dp[0] (pick 3, left tree have 2 nodes, right tree empty)
	// dp[4]
	//     = dp[0] * dp[3]
	//     + dp[1] * dp[2]
	//     + dp[2] * dp[1]
	//     + dp[3] * dp[0]
	dp := make(map[int]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
