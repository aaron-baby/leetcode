package main

func isValid(s string) bool {
	var x rune
	var stack []rune
	leftPair := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else if c == ')' || c == ']' || c == '}' {
			// no elements to pop
			if len(stack) == 0 {
				return false
			}
			// pop stack
			x, stack = stack[len(stack)-1], stack[:len(stack)-1]
			if x != leftPair[c] {
				return false
			}
		}
	}
	return len(stack) == 0
}
