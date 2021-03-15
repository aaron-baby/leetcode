package stack

import (
	"strconv"
)

func calculate(s string) int {
	var ns string
	var res int
	var sign = 1
	var num int
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			ns += string(c)
		}
		if c == '(' {
			start := i
			cnt := 0
			for ; i < len(s); i++ {
				if s[i] == '(' {
					cnt++
				}
				if s[i] == ')' {
					cnt--
				}
				if cnt == 0 {
					break
				}
			}
			num = calculate(s[start+1 : i])
		}
		if c == '+' || c == '-' || i == len(s)-1 {
			if len(ns) > 0 {
				num, _ = strconv.Atoi(ns)
			}
			res += sign * num
			sign = 1
			if '-' == c {
				sign = -1
			}
			ns = ""
		}
	}

	return res
}
