package leetcode_0097

// https://leetcode.cn/problems/interleaving-string/?envType=study-plan-v2&envId=top-interview-150
func isInterleave(s1 string, s2 string, s3 string) bool {
	var (
		n1, n2, n3 = len(s1), len(s2), len(s3)
	)
	// param check
	if n1+n2 != n3 {
		return false
	}

	dp := make([][]bool, n1+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n2+1)
	}
	// 0. 状态定义
	// dp[i][j]: s1[:i] 与 s2[:j]是否可以交错成字符串 s3[:i+j]
	// 1. init state
	dp[0][0] = true
	// 1. 只用s2
	for i := 1; i <= n1; i++ {
		dp[0][i] = dp[0][i-1] && s2[i-1] == s3[i-1]
	}
	// 2. 只用s1
	for i := 1; i <= n1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}
	for i := 1; i <= n1; i++ {
		for j := 1; j <= n2; j++ {
			p := i + j - 1
			dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[p])
			dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[p])
		}
	}
	return dp[n1][n2]
}
