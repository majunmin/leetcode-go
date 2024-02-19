package lintcode_0440

func BackPackIII(a []int, v []int, m int) int {
	return dpSolution2(a, v, m)
}

// 空间优化的动态规划解法
// 因为 选择前i个物品的最大价值与只与选择前i-1个物品相关
func dpSolution2(a []int, v []int, m int) int {
	n := len(a)
	if n != len(v) {
		panic("invalid param")
	}
	if n == 0 {
		return 0
	}
	dp := make([]int, m+1)
	//dp[0] = 0
	for i := 0; i < n; i++ {
		for j := a[i]; j <= m; j++ {
			dp[j] = max(dp[j], dp[j-a[i]]+v[i])
		}
	}
	return dp[m]
}

func dpSolution1(a []int, v []int, m int) int {
	// write your code here
	// param check
	n := len(a)
	//
	if n != len(v) {
		panic("invalid param, len(a) != len(v)")
	}
	if n == 0 {
		return 0
	}
	// 状态定义
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// init state

	// 递推公式
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			for k := 1; k*a[i-1] <= j; k++ {
				dp[i][j] = max(dp[i][j], dp[i-1][j-k*a[i-1]]+k*v[i-1])
			}
		}
	}
	return dp[n][m]
}
