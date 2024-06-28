package leetcode_100280

import "sort"

// https://leetcode.cn/contest/biweekly-contest-128/problems/minimum-rectangles-to-cover-points/
func minRectanglesToCoverPoints(points [][]int, w int) int {
	var (
		pointx   = make([]int, 0, len(points))
		pointSet = make(map[int]bool)
		result   int
	)
	for _, p := range points {
		if pointSet[p[0]] {
			continue
		}
		pointSet[p[0]] = true
		pointx = append(pointx, p[0])
	}
	// 排序
	sort.Ints(pointx)
	if w == 0 {
		return len(pointx)
	}
	end := -1
	for i := 0; i < len(points); i++ {
		if pointx[i] <= end {
			continue
		}
		end = pointx[i] + w
		result++
	}
	return result
}
