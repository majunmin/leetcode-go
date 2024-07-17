package leetcode_0221

// https://leetcode.cn/problems/maximal-square/?envType=study-plan-v2&envId=top-interview-150
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	// 状态定义
	// dp[i][j] 以 matrix[i-1][j-1] 为右下角的最大正方形边长
	// 最短的行 or 最短的列
	// 状态转移
	// dp[i][j] = 0 if matrix[i-1][j-1] == 0
	// dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])+ 1 if matrix[i-1][j-1] == 1
	var (
		m, n = len(matrix), len(matrix[0])
		dp   = make([][]int, m+1)
	)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if matrix[i-1][j-1] == 0 {
				continue
			}
			dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
		}
	}
	return dp[m][n]
}
