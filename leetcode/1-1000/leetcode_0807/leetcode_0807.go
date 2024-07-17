package leetcode_0807

// https://leetcode.cn/problems/max-increase-to-keep-city-skyline/?envType=daily-question&envId=2024-07-14
func maxIncreaseKeepingSkyline(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		m, n   = len(grid), len(grid[0])
		rowMax = make([]int, m)
		colMax = make([]int, n)
		res    int
	)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			rowMax[i] = max(rowMax[i], grid[i][j])
			colMax[j] = max(colMax[j], grid[i][j])
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res += min(rowMax[i], colMax[j]) - grid[i][j]
		}
	}
	return res
}
