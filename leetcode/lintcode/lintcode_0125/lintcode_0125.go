package lintcode_0125

// https://www.lintcode.com/problem/125/?showListFe=true&page=1&submissionStatus=ACCEPTED&pageSize=50
func BackPackII(m int, a []int, v []int) int {
	return dpSolution2(m, a, v)
}

// 空间优化的动态规划解法
// 因为 选择前i个物品的最大价值与只与选择前i-1个物品相关
func dpSolution2(m int, a []int, v []int) int {
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
		for j := m; j >= a[i]; j-- {
			dp[j] = max(dp[j], dp[j-a[i]]+v[i])
		}
	}
	return dp[m]
}

func dpSolution1(m int, a []int, v []int) int {
	// write your code here
	// param check
	n := len(a)
	if n != len(v) {
		panic("invalid param")
	}
	if n == 0 {
		return 0
	}

	//
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}
	// 1. state define
	// dp[i][j] :背包容量为 i, 放入前i件物品的最大价值v
	// 2. init state
	//  dp[0][0] = 0
	// 3. 递推公式
	// dp[i][j] = max(dp[i-1][j], dp[i-1][j-a[i-1]] + v[i-1])
	for i := 1; i <= n; i++ {
		for j := 0; j <= m; j++ {
			dp[i][j] = dp[i-1][j]
			if j > a[i-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-a[i-1]]+v[i-1])
			}
		}
	}
	return dp[n][m]
}
