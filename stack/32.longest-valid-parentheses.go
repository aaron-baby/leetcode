package stack

import (
	"gitlab.com/aaw/leetcode/lib"
)

func longestValidParentheses(s string) int {
	res, start, n := 0, 0, len(s)
	var stack []int
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else if s[i] == ')' {
			if len(stack) == 0 {
				start = i + 1
			} else {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					res = lib.Max(res, i-stack[len(stack)-1])
				} else {
					res = lib.Max(res, i-start+1)
				}
			}
		}
	}
	return res
}
