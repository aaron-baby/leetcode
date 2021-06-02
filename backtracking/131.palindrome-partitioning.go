package backtracking

func partition(s string) [][]string {
	var res = [][]string{}
	c := []string{}
	findSub(s, c, &res)
	return res
}
func findSub(remain string, c []string, res *[][]string) {
	if remain == "" {
		b := make([]string, len(c))
		copy(b, c)
		*res = append(*res, b)
		return
	}
	for i := 0; i < len(remain); i++ {
		if !isPalindrome(remain[0 : i+1]) {
			continue
		}
		c = append(c, remain[0:i+1])
		findSub(remain[i+1:], c, res)
		c = c[:len(c)-1]
	}
}
func isPalindrome(s string) bool {
	var i, j = 0, len(s) - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
