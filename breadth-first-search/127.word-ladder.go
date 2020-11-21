package breadth_first_search

func ladderLength(beginWord string, endWord string, wordList []string) int {
	var d = make(map[string]struct{})
	for _, word := range wordList {
		d[word] = struct{}{}
	}
	var q = []string{beginWord}
	var w string
	var pathCnt = make(map[string]int)
	pathCnt[beginWord] = 1
	for len(q) > 0 {
		w, q = q[0], q[1:]
		for i := 0; i < len(w); i++ {
			var ch byte
			for ch = 'a'; ch <= 'z'; ch++ {
				if ch == w[i] {
					continue
				}
				nw := w[:i] + string(ch) + w[i+1:]
				_, ok := d[nw]
				if ok && nw == endWord {
					return pathCnt[w] + 1
				}
				_, pok := pathCnt[nw]
				if ok && !pok {
					q = append(q, nw)
					pathCnt[nw] = pathCnt[w] + 1
				}
			}
		}
	}
	return 0
}
