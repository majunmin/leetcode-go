package leetcode_0056

import "slices"

// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked
func merge(intervals [][]int) [][]int {
	// param check
	var result [][]int
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := 0; i < len(intervals); i++ {
		curInterval := intervals[i]
		if len(result) == 0 {
			result = append(result, intervals[i])
			continue
		}
		if curInterval[0] <= result[len(result)-1][1] {
			result[len(result)-1][1] = max(result[len(result)-1][1], curInterval[1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
