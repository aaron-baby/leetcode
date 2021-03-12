package greedy_algorithm

func removeDuplicateLetters(s string) string {
	var counter = make(map[byte]int, 26)
	var visited = make(map[byte]bool, 26)
	for i := 0; i < len(s); i++ {
		counter[s[i]] += 1
	}
	var st []byte
	for i := 0; i < len(s); i++ {
		c := s[i]
		counter[c]--
		if visited[c] {
			continue
		}
		for len(st) != 0 && c < st[len(st)-1] && counter[st[len(st)-1]] != 0 {
			top := st[len(st)-1]
			st = st[:len(st)-1]
			visited[top] = false
		}
		st = append(st, c)
		visited[c] = true
	}
	return string(st)
}
