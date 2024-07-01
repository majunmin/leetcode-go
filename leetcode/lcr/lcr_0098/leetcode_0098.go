package lcr_0098

// https://leetcode.cn/problems/2AoeFn/
func uniquePaths(m int, n int) int {
	// 动态规划的解法
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

// 不同路径2 带障碍物
// https://leetcode.cn/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	var (
		m = len(obstacleGrid)
		n = len(obstacleGrid[0])
	)
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	// 动态规划的解法
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	// init state
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = dp[i-1][0]
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}
