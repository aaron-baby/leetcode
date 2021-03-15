package stack

func calculate_ii(s string) int {
	var st []int
	num := 0
	//operator before number
	var sign byte
	sign = '+'
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = 10*num + int(c-'0')
		}
		if c == '-' || c == '+' || c == '*' || c == '/' || i == len(s)-1 {
			if sign == '-' || sign == '+' {
				if sign == '-' {
					num = -num
				}
				st = append(st, num)
			}
			if sign == '*' {
				top := st[len(st)-1]
				st = st[:len(st)-1]
				st = append(st, top*num)
			}
			if sign == '/' {
				top := st[len(st)-1]
				st = st[:len(st)-1]
				st = append(st, top/num)
			}
			sign = c
			num = 0
		}
	}
	var res int
	for len(st) > 0 {
		res += st[len(st)-1]
		st = st[:len(st)-1]
	}
	return res
}
