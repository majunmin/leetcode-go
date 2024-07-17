package leetcode_3212

// https://leetcode.cn/problems/count-submatrices-with-equal-frequency-of-x-and-y/
func numberOfSubmatrices(grid [][]byte) int {
	// 动态规划
	// dp[i][j][0] :矩形 以 [0,0] - [i,j] 之间的矩形的 X频数
	// dp[i][j][1] :矩形 以 [0,0] - [i,j] 之间的矩形的 Y频数
	var (
		m, n   = len(grid), len(grid[0])
		dp     = make([][][2]int, m+1)
		result int
	)
	for i := 0; i <= m; i++ {
		dp[i] = make([][2]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if grid[i-1][j-1] == 'X' {
				dp[i][j][0] = 1
			} else if grid[i-1][j-1] == 'Y' {
				dp[i][j][1] = 1
			}
			dp[i][j][0] += dp[i-1][j][0] + dp[i][j-1][0] - dp[i-1][j-1][0]
			dp[i][j][1] += dp[i-1][j][1] + dp[i][j-1][1] - dp[i-1][j-1][1]
			if dp[i][j][0] > 0 && dp[i][j][0] == dp[i][j][1] {
				result++
			}
		}
	}
	return result
}
