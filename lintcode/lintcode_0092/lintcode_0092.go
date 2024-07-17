package lintcode_0092

// https://www.lintcode.com/problem/92/?showListFe=true&page=1&submissionStatus=ACCEPTED&pageSize=50
func BackPack(m int, a []int) int {
	// write your code here
	// param check
	n := len(a)
	if n == 0 {
		return 0
	}
	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, m+1)
	}
	// init state
	dp[0][0] = true
	// 递推公式
	for i := 1; i <= n; i++ {
		// 不选择当期那物品
		for j := 0; j <= m; j++ {
			// 不选当前物品  || 选择当前物品
			dp[i][j] = dp[i][j] || (j >= a[i-1] && dp[i-1][j-a[i-1]])
		}
	}
	for i := m; i >= 0; i-- {
		if dp[n][i] {
			return i
		}
	}
	return 0
}
