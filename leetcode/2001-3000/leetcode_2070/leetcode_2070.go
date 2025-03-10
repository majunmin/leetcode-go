package leetcode_2070

import "slices"

// https://leetcode.cn/problems/most-beautiful-item-for-each-query/?envType=daily-question&envId=2025-03-10
func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int {
		if a[0] != b[0] {
			return a[0] - b[0]
		}
		return b[1] - a[1]
	})
	var (
		preMax = make([]int, 0, len(queries))
		result = make([]int, 0, len(queries))
		curMax = 0
	)
	for _, item := range items {
		curMax = max(curMax, item[1])
		preMax = append(preMax, curMax)
	}
	for _, target := range queries {
		idx := upperBound(items, target)
		if idx == -1 {
			result = append(result, 0)
		} else {
			result = append(result, preMax[idx])
		}
	}
	return result
}

func upperBound(items [][]int, target int) int {
	l, r := -1, len(items)
	for l+1 < r {
		mid := l + (r-l)/2
		if items[mid][0] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// queries,
