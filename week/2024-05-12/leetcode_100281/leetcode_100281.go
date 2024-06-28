package leetcode_100281

import "math"

// https://leetcode.cn/contest/weekly-contest-397/problems/maximum-difference-score-in-a-grid/
func maxScore(grid [][]int) int {
	// param check
	var (
		result = math.MinInt
		m, n   = len(grid), len(grid[0])
		dp     = make([][]int, m)
	)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = -1000000
		}
	}
	// dp[i][j] :走到单元格 grid[i][j]的maxScore
	// init state
	// dp[0][0] = 0
	// dp[i][j] = max()
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			// 按行
			k := i - 1
			for k >= 0 {
				dp[i][j] = max(dp[i][j], grid[i][j]-grid[k][j], grid[i][j]-grid[k][j]+dp[k][j])
				k--
			}
			// 按列统计
			k = j - 1
			for k >= 0 {
				dp[i][j] = max(dp[i][j], grid[i][j]-grid[i][k], grid[i][j]-grid[i][k]+dp[i][k])
				k--
			}
			result = max(result, dp[i][j])
		}
	}
	return result
}
