package leetcode_0986

// https://leetcode.cn/problems/interval-list-intersections/
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	// param check
	firstLength, secondLength := len(firstList), len(secondList)
	if firstLength == 0 || secondLength == 0 {
		return nil
	}
	var result [][]int
	firstIdx, secondIdx := 0, 0
	for firstIdx < firstLength && secondIdx < secondLength {
		interval1, interval2 := firstList[firstIdx], secondList[secondIdx]
		if interval1[1] < interval2[0] || interval2[1] < interval1[0] {
			continue
		}
		// 存在交集
		start, end := max(interval1[0], interval2[0]), min(interval1[1], interval2[1])
		result = append(result, []int{start, end})
		if interval1[1] < interval2[1] {
			firstIdx++
		} else if interval1[1] > interval2[1] {
			secondIdx++
		} else {
			firstIdx++
			secondIdx++
		}
	}
	return result
}
