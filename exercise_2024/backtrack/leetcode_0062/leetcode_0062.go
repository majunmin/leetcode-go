package leetcode_0062

// https://leetcode.cn/problems/unique-paths/?envType=study-plan-v2&envId=top-100-liked
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n)
	}

	// init
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

func uniquePaths2(m int, n int) int {
	dp := make([]int, n)

	// init
	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if i == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}
