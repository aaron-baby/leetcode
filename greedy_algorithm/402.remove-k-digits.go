package greedy_algorithm

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var st []byte
	keep := len(num) - k
	for i := 0; i < len(num); i++ {
		// pop top elements of stack
		for k > 0 && len(st) > 0 && st[len(st)-1] > num[i] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, num[i])
	}
	st = st[:keep]
	// remove leading zeros, exclude last digit
	j := 0
	for j < len(st)-1 && st[j] == '0' {
		st = st[1:]
	}
	return string(st)
}
