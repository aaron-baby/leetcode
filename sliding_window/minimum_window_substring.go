package sliding_window

// 76. Minimum Window Substring
// https://leetcode.com/problems/minimum-window-substring/
func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 || len(t) > len(s) {
		return ""
	}
	l, r := 0, 0
	counter, requirement := make(map[uint8]int), make(map[byte]int)
	for _, c := range []byte(t) {
		counter[c] = 0
		requirement[c] = 0
	}
	for i := range t {
		ch := t[i]
		requirement[ch] += 1
	}
	// answer's length, left cursor, right cursor
	ans := [3]int{-1, l, r}
	for r < len(s) {
		_, ok := counter[s[r]]
		if ok {
			counter[s[r]] += 1
		}
		r++
		if meet_criterion(counter, requirement) {
			for l <= r {
				if ans[0] == -1 || r-l+1 < ans[0] {
					ans = [3]int{r - l + 1, l, r}
				}
				// try to shrink window sizing by move left cursor forward
				_, ok := counter[s[l]]
				if ok {
					counter[s[l]] -= 1
				}
				l++
				if !meet_criterion(counter, requirement) {
					break
				}
			}
		}
	}
	if ans[0] == -1 {
		return ""
	}
	return s[ans[1]:ans[2]]
}

func meet_criterion(counter map[uint8]int, requirement map[byte]int) bool {
	for c, v := range counter {
		if v < requirement[c] {
			return false
		}
	}
	return true
}
