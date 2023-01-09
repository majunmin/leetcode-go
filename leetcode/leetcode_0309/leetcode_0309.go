package leetcode_0309

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}
	// dp[i][0]: 第 i 天 结束时, 持有股票
	// dp[i][1]: 第 i 天 结束时, 不持有股票,处于冷冻期 🥶
	// dp[i][2]: 第 i 天 结束时, 不持有股票, 不处于冷冻期
	dp := make([][3]int, n)
	// init State
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0

	// 状态转移方程
	// dp[i][0] = max(dp[i-1][0], dp[i-1][2] - prices[i])
	// dp[i][1] = dp[i-1][0] + prices[i]
	// dp[i][2] = max(dp[i-1][1],  dp[i-1][2])

	for i := 1; i < n; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = dp[i-1][0] + prices[i]
		dp[i][2] = maxInt(dp[i-1][2], dp[i-1][1])
	}
	return maxInt(dp[n-1][1], dp[n-1][2])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
