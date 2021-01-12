package breadth_first_search

import "container/list"

func openLock(deadends []string, target string) int {
	queue := list.New()
	visited := make(map[string]bool)
	deads := make(map[string]bool)
	for _, s := range deadends {
		deads[s] = true
	}
	var step int
	queue.PushBack("0000")
	visited["0000"] = true
	// 广度优先
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			cur := dequeue(queue).(string)
			if _, ok := deads[cur]; ok {
				continue
			}
			if cur == target {
				return step
			}
			for j := 0; j < 4; j++ {
				next := getNext(cur, j)
				if _, ok := visited[next]; !ok {
					queue.PushBack(next)
					visited[next] = true
				}

				prev := getPrev(cur, j)
				if _, ok := visited[prev]; !ok {
					queue.PushBack(prev)
					visited[prev] = true
				}
			}
		}
		step++
	}
	return -1
}
func getNext(s string, i int) string {
	if s[i] == '9' {
		return s[:i] + "0" + s[i+1:]
	}
	return s[:i] + string(s[i]+1) + s[i+1:]
}
func getPrev(s string, i int) string {
	if s[i] == '0' {
		return s[:i] + "9" + s[i+1:]
	}
	return s[:i] + string(s[i]-1) + s[i+1:]
}
func dequeue(queue *list.List) interface{} {
	if queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		return e.Value
	}
	return nil
}
