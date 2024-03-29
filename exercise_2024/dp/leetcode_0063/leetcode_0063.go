package leetcode_0063

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// param check
	if len(obstacleGrid) == 0 || len(obstacleGrid[0]) == 0 {
		return 0
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	rowSize := len(obstacleGrid)
	colSize := len(obstacleGrid[0])
	// status define and init
	dp := make([][]int, rowSize)
	for i := 0; i < rowSize; i++ {
		dp[i] = make([]int, colSize)
	}
	//
	dp[0][0] = 1
	for i := 1; i < rowSize; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < colSize; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}
	// 递推公式
	for i := 1; i < rowSize; i++ {
		for j := 1; j < colSize; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[rowSize][colSize]
}
