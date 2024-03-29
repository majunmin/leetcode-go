package leetcode_0309

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
func maxProfit(prices []int) int {
	dp := make([][3]int, len(prices))
	dp[0][0] = 0          // 非处于冷冻期持有现金的最大价值
	dp[0][1] = -prices[0] // 持有股票的最大价值
	dp[0][2] = -prices[0] // 持有现金处于冷冻期的最大价值
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return max(dp[len(dp)-1][0], dp[len(dp)-1][2])
}
