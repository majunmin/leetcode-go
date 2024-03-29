package leetcode_0056

import "sort"

func merge(intervals [][]int) [][]int {
	//
	var result [][]int
	if len(intervals) < 2 {
		return intervals
	}

	// sort intervals
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if prev[1] < intervals[i][0] {
			result = append(result, prev)
			prev = intervals[i]
			continue
		}
		if prev[1] >= intervals[i][1] {
			continue
		}
		prev[1] = intervals[i][1]
	}
	result = append(result, prev)
	//
	return result
}
