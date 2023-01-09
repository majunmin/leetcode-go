package leetcode_1620

import "math"

// https://leetcode.cn/problems/coordinate-with-maximum-network-quality/
func bestCoordinate(towers [][]int, radius int) []int {
	// param check
	if len(towers) == 0 {
		return nil
	}
	//
	xmin, xmax := 0, 50
	ymin, ymax := 0, 50
	cx, cy, cq := 0, 0, 0

	//  from bottom to top, from left to right
	for i := xmin; i <= xmax; i++ {
		for j := ymin; j <= ymax; j++ {
			// 计算坐标 i，j 的信号
			quality := 0
			for _, item := range towers {
				d2 := (item[0]-i)*(item[0]-i) + (item[1]-j)*(item[1]-j)
				// 理解题目的边界条件 (以内...)
				if d2 <= radius*radius {
					quality += int(float64(item[2]) / (1 + math.Sqrt(float64(d2))))
				}
			}
			if quality > cq {
				cx, cy, cq = i, j, quality
			}
		}
	}

	return []int{cx, cy}
}
