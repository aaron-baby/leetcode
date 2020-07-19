package sort

import (
	"gitlab.com/aaw/leetcode/lib"
	"sort"
)

// 56. Merge Intervals
// https://leetcode.com/problems/merge-intervals/
func merge(intervals [][]int) [][]int {
	if len(intervals) < 1 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[j][0] > intervals[i][0]
	})
	// last merged index
	var index = 0
	for i := 0; i < len(intervals); i++ {
		if intervals[index][1] >= intervals[i][0] {
			intervals[index][0] = lib.Min(intervals[index][0], intervals[i][0])
			intervals[index][1] = lib.Max(intervals[index][1], intervals[i][1])
		} else {
			index++
			intervals[index] = intervals[i]
		}
	}

	return intervals[:index+1]
}
