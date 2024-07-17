package lcr_097

// https://leetcode.cn/problems/21dk04/
func numDistinct(s string, t string) int {
	var (
		m, n = len(s), len(t)
	)

	if m < n {
		return 0
	}

	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j == 0 {
				dp[i][j] = 1
				continue
			}
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[m][n]
}
