package offer_014

// https://leetcode.cn/problems/jian-sheng-zi-lcof/?envType=study-plan&id=lcof
func cuttingRope(n int) int {
	// param check
	if n <= 1 {
		return -1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = maxInt(dp[i], maxInt((i-j)*j, dp[i-j]*j))
		}
	}
	return dp[n]

}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
