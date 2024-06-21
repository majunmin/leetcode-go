package leetcode_2813

import "slices"

// https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence/?envType=daily-question&envId=2024-06-13
//func findMaximumElegance(items [][]int, k int) int64 {
//	slices.SortFunc(items, func(a, b []int) int {
//
//	})
//}

func merge(intervals [][]int) [][]int {
	// sort by interval[0] asc
	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})
	var (
		result = make([][]int, 0)
	)
	result = append(result, intervals[0])
	for i := 1; i < len(intervals); i++ {
		if result[len(result)-1][1] < intervals[i][0] {
			result = append(result, intervals[i])
			continue
		}
		// merge
		result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
	}

	return result
}

//
