package leetcode_0062

// https://leetcode.cn/problems/unique-paths/
func uniquePaths(m int, n int) int {
	return dfsSolution(m, n)
}

// 该方法 会导致超时, 会有大量重复计算,可以再次优化为记忆化递归,也就是类似于递归解法了
func dfsSolution(m int, n int) int {
	if m == 0 || n == 0 {
		return 1
	}
	return dfsSolution(m-1, n) + dfsSolution(m, n-1)
}

// 动态规划 空间优化的解法
func solution2(m int, n int) int {
	if m <= 0 || n <= 0 {
		panic("m and n must be greater than 0")
	}
	// param check
	if m < 2 || n < 2 {
		return 1
	}
	// m n
	dp := make([]int, n)
	// init
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j] + dp[j-1]
		}
	}
	return dp[n-1]
}

// 表格法 动态规划
func solution1(m int, n int) int {
	// param invalid
	if m < 1 || n < 1 {
		panic("m and n must be greater than 0")
	}
	if m < 2 || n < 2 {
		return 1
	}
	dp := make([][]int, 0, m)
	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}
	// 状态定义
	// dp[i][j]: i * j 网格, 从 左上角的到右下角的 路径数
	// 这个问题有点类似 青蛙跳台 和   爬楼梯问题
	// dp[i][j] = 1 (i == 0 || j == 0)
	// dp[i][j] = dp[i-1][j] + dp[i][j-1]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// init state
			if i == 0 || j == 0 {
				dp[i][j] = 1
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
