package lcr_074

import "slices"

// https://leetcode.cn/problems/SsGoHC/
func merge(intervals [][]int) [][]int {
	// param check
	if len(intervals) <= 1 {
		return intervals
	}
	//按照区间 start 排序从小至大
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var (
		result = make([][]int, 0)
	)
	result = append(result, intervals[0])
	for _, interval := range intervals {
		if interval[0] > result[len(result)-1][1] {
			result = append(result, interval)
			continue
		}
		// merge
		result[len(result)-1][1] = max(result[len(result)-1][1], interval[1])
	}
	return result
}
