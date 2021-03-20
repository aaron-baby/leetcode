package greedy_algorithm

func numMatchingSubseq(s string, words []string) int {
	var dict = make(map[byte][]int)
	// initial s letter occurrence
	for i := 0; i < len(s); i++ {
		if _, ok := dict[s[i]]; ok {
			dict[s[i]] = append(dict[s[i]], i)
		} else {
			dict[s[i]] = []int{i}
		}
	}
	var res int
	for _, v := range words {
		if isMatch(v, 0, 0, dict) {
			res += 1
		}
	}
	return res
}
func isMatch(w string, wi, di int, d map[byte][]int) bool {
	if wi == len(w) {
		return true
	}
	s, present := d[w[wi]]
	if !present || di > s[len(s)-1] {
		return false
	}
	var i int
	for _, v := range s {
		if v >= di {
			i = v
			break
		}
	}
	return isMatch(w, wi+1, i+1, d)
}
