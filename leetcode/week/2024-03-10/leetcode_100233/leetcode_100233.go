package leetcode_100233

import "slices"

// https://leetcode.cn/contest/weekly-contest-388/problems/apple-redistribution-into-boxes/
func minimumBoxes(apple []int, capacity []int) int {
	var total int
	for _, w := range apple {
		total += w
	}
	if total == 0 {
		return 0
	}
	//贪心, 箱子从大到小排序
	slices.SortFunc(capacity, func(a, b int) int {
		return b - a
	})
	//
	var (
		cnt    int
		totalC int
	)
	for _, c := range capacity {
		totalC += c
		cnt++
		if totalC >= total {
			return cnt
		}
	}
	return -1
}
