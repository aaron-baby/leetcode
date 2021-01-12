package breadth_first_search

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegrees := make(map[int]int)
	adjacencies := map[int][]int{}
	for _, edge := range prerequisites {
		start, end := edge[1], edge[0]
		indegrees[end]++
		adjacencies[start] = append(adjacencies[start], end)
	}
	// Create a queue and enqueue all vertices with
	// indegree 0
	var q []int
	for i := 0; i < numCourses; i++ {
		if indegrees[i] == 0 {
			q = append(q, i)
		}
	}

	var course, visited int
	for len(q) > 0 {
		course, q = q[len(q)-1], q[:len(q)-1]
		for _, neighbour := range adjacencies[course] {
			indegrees[neighbour]--
			if indegrees[neighbour] == 0 {
				q = append(q, neighbour)
			}
		}
		visited++
	}

	return visited == numCourses
}
