package leetcode_0010

// https://leetcode.cn/problems/regular-expression-matching/description/
func isMatch(s string, p string) bool {
	// 动态规划, 这个问题 要考虑的场景挺多的,需要想清楚,写明白,下次可以更好地回忆起来.

	// param check
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	// init state
	dp[0][0] = true
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	// 状态转移方程
	//dp[i][j] = dp[i-1][j-1]   s[i] == p[j]
	//dp[i][j] = dp[i-1][j-1]   p[j] = '.'
	// p[j]  = '*'
	// - dp[i][j] = dp[i][j-1]  // ###a  ###a*
	// - dp[i][j] = dp[i][j-2]  // ###  ###a*
	// - dp[i][j] = dp[i-1][j]  // ###aaaaaa  ###a*
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == s[i-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				//p[i] == '*'
				if p[j-2] != s[i-1] && p[j-2] != '.' {
					dp[i][j] = dp[i][j-2]
				} else {
					dp[i][j] = dp[i][j-1] || dp[i][j-2] || dp[i-1][j]
				}
			}
		}
	}
	return dp[m][n]
}
