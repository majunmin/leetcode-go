package leetcode_0064

// https://leetcode.cn/problems/minimum-path-sum/?envType=study-plan-v2&envId=top-100-liked
func minPathSum(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}
	// process
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i == 0 {
				grid[i][j] += grid[i][j-1]
				continue
			}
			if j == 0 {
				grid[i][j] += grid[i-1][j]
				continue
			}
			grid[i][j] += min(grid[i-1][j], grid[i][j-1])
		}
	}
	return grid[m-1][n-1]
}
