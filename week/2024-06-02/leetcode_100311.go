package _024_06_02

import (
	"slices"
)

// https://leetcode.cn/problems/count-days-without-meetings/
func countDays(days int, meetings [][]int) int {
	var (
		meetDays   int
		start, end = 1, 0
	)
	slices.SortFunc(meetings, func(a, b []int) int {
		return a[0] - b[0]
	})
	for i := 1; i < len(meetings); i++ {
		meet := meetings[i]
		if meet[0] <= end {
			start = min(start, meet[0])
			end = max(end, meet[1])
		} else {
			meetDays += end - start + 1
		}
	}
	// 统计最后一个区间
	meetDays += end - start + 1
	return days - meetDays
}
