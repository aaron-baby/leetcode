package sort

import "gitlab.com/aaw/leetcode/lib"

func insert(intervals [][]int, newInterval []int) [][]int {
	n, cur := len(intervals), 0
	res := make([][]int, 0)
	for cur < n && intervals[cur][1] < newInterval[0] {
		res = append(res, intervals[cur])
		cur++
	}
	for cur < n && intervals[cur][0] <= newInterval[1] {
		newInterval[0] = lib.Min(intervals[cur][0], newInterval[0])
		newInterval[1] = lib.Max(intervals[cur][1], newInterval[1])
		cur++
	}
	res = append(res, newInterval)
	for ; cur < n; cur++ {
		res = append(res, intervals[cur])
	}
	return res
}
