package exercise_2024

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	// dp[i][0]: 持有现金的最大收益
	// dp[i][1]: 持有股票的最大收益
	stock, crash := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		crash = max(crash, crash+prices[i])
		stock = max(stock, -prices[i])
	}
	return crash
}
