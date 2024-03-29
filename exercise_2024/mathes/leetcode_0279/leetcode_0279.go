package leetcode_0279

// https://leetcode.cn/problems/perfect-squares/?envType=study-plan-v2&envId=top-100-liked
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		// 最坏的情况,都是1 组成
		dp[i] = i
		for j := 1; i-j*j > 0; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
