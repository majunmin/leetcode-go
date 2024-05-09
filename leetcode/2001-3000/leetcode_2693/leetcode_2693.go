package leetcode_2693

// https://leetcode.cn/problems/find-the-width-of-columns-of-a-grid/
func findColumnWidth(grid [][]int) []int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return nil
	}
	var (
		m, n   = len(grid), len(grid[0])
		result = make([]int, n)
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j] = max(result[j], calLen(grid[i][j]))
		}
	}
	return result
}

func calLen(i int) int {
	var length int
	if i < 0 {
		length += 1
		i = -i
	}
	for i > 0 {
		length += 1
		i /= 10
	}
	return length
}
