package exercise_2024

import "slices"

// https://leetcode.cn/problems/merge-intervals/?envType=study-plan-v2&envId=top-100-liked
func merge(intervals [][]int) [][]int {
	// param check
	if len(intervals) == 0 {
		return nil
	}
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var result [][]int
	prev := intervals[0]
	for i := 1; i < len(intervals); i++ {
		cur := intervals[i]
		if prev[1] < cur[0] {
			result = append(result, prev)
			prev = cur
			continue
		}
		prev = []int{prev[0], max(prev[1], cur[1])}
	}
	result = append(result, prev)
	return result
}
