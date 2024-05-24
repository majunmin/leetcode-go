package leetcode_0452

import "slices"

// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/description/?envType=study-plan-v2&envId=top-interview-150
func findMinArrowShots(points [][]int) int {
	// param check
	if len(points) == 0 {
		return 0
	}

	// 按照区间右端点 升序排序
	slices.SortFunc(points, func(a, b []int) int {
		return a[1] - b[1]
	})
	var (
		cnt = 1
		end = points[0][1]
	)
	for i := 1; i < len(points); i++ {
		if points[i][0] <= end {
			continue
		}
		cnt++
		end = points[i][1]
	}
	return cnt
}
