package main

func countWays(M int, N []int) int {
	var (
		dp     = make([][]int, len(N))
		minVal = N[0]
	)
	for i := 0; i < len(N); i++ {
		dp[i] = make([]int, M+1)
	}
	dp[0][0] = 1
	for i := 0; i < len(dp); i++ {
		dp[i][0] = 1
		for j := 0; j <= M; j++ {
			if i > 0 {
				dp[i][j] = dp[i-1][j]
			}
			if j-N[i] >= 0 {
				dp[i][j] += dp[i][j-N[i]]
			}
		}
	}
	// 考虑 N 之外的数
	var result int
	for i := 0; i < minVal; i++ {
		result += dp[len(N)-1][M-i]
	}
	return result
}

func countWays2(M int, N []int) int {
	var (
		dp     = make([]int, M+1)
		minVal = N[0]
	)
	dp[0] = 1
	for _, num := range N {
		for i := M - 1; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}

	// 考虑 N 之外的数
	var result int
	for i := 0; i < minVal; i++ {
		result += dp[M-N[i]]
	}
	return result
}
