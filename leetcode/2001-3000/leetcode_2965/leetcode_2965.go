package leetcode_2965

// https://leetcode.cn/problems/find-missing-and-repeated-values/
func findMissingAndRepeatedValues(grid [][]int) []int {
	var (
		n    = len(grid)
		cnts = make([]int, n*n+1)
	)
	for _, row := range grid {
		for _, num := range row {
			cnts[num]++
		}
	}
	var repeat, miss int
	for i, cnt := range cnts {
		if cnt == 0 {
			miss = i
		}
		if cnt == 2 {
			repeat = i
		}
	}
	return []int{repeat, miss}
}
