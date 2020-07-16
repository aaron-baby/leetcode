package main

import (
	"sort"
)

// 406. Queue Reconstruction by Height
func reconstructQueue(people [][]int) [][]int {
	// sort by height descend, while height is same, let people who have less k value on the front
	sort.Slice(people, func(i, j int) bool {
		return people[i][0] > people[j][0] || (people[i][0] == people[j][0] && people[i][1] < people[j][1])
	})
	for i, p := range people {
		copy(people[p[1]+1:i+1], people[p[1]:i])
		people[p[1]] = p
	}

	return people
}
