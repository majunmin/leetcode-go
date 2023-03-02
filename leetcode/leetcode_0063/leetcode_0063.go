package leetcode_0063

// https://leetcode.cn/problems/unique-paths-ii/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	return solution2(obstacleGrid)
}

// 动态规划2 空间优化
func solution2(obstacleGrid [][]int) int {
	// param check
	if len(obstacleGrid) <= 0 || len(obstacleGrid[0]) <= 0 {
		panic("obstacleGrid should not be empty.")
	}
	// 一开始就是石头, 那就直接返回 0
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([]int, n)
	// init state
	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 0; j < n; j++ {
			// 竖向遇到第一个石头, dp[0] = 0, 否则不变
			if obstacleGrid[i][j] == 1 {
				dp[j] = 0
				continue
			}
			if j == 0 {
				// dp[j] = dp[j]
				continue
			}

			dp[j] = dp[j] + dp[j-1]
		}
	}

	return dp[len(dp)-1]
}

// 动态规划
func solution1(obstacleGrid [][]int) int {
	// param check
	if len(obstacleGrid) <= 0 || len(obstacleGrid[0]) <= 0 {
		panic("obstacleGrid should not be empty.")
	}
	// 一开始就是石头, 那就直接返回 0
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	m, n := len(obstacleGrid), len(obstacleGrid[0])

	// dp[i][j] :从 [0,0] 到 [i, j] 的路径数
	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	// init state
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
