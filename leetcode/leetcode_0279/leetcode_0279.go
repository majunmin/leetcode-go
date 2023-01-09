package leetcode_0279

func numSquares(n int) int {
	// param check
	if n <= 0 {
		return 0
	}

	//dp[i] = dp[i - j*j ] + 1
	dp := make([]int, n+1)
	// init state
	dp[0] = 0
	for i := 1; i <= n; i++ {
		// 最坏的情况就是 每次都比前一次 + 1(为了 与 0 比较,初始化一个 比较大的值)
		dp[i] = dp[i-1] + 1
		for j := 1; i-j*j >= 0; j++ {
			dp[i] = minInt(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
