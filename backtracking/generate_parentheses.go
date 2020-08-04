package backtracking

func generateParenthesis(n int) []string {
	var ans []string
	addParentheses(&ans, "", 0, 0, n)

	return ans
}

func addParentheses(ans *[]string, cur string, open, close, max int) {
	// exit condition
	if len(cur) == max*2 {
		output := *ans
		*ans = append(output, cur)
		return
	}
	if open < max {
		addParentheses(ans, cur+"(", open+1, close, max)
	}
	if close < open {
		addParentheses(ans, cur+")", open, close+1, max)
	}
}
