package leetcode_0123

import "math"

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	//
	dp := make([][3][2]int, len(prices))
	//dp[i][j][k]
	// j：0,1,2 : 表示第j次交易
	// j:0,1 : 表示持有现金/股票的最大收益
	dp[0][0][1] = math.MinInt
	dp[0][1][1] = -prices[0]
	dp[0][2][1] = math.MinInt
	for i := 1; i < len(prices); i++ {
		dp[i][1][0] = max(dp[i-1][1][0], dp[i-1][1][1]+prices[i])
		dp[i][1][1] = max(dp[i-1][1][1], dp[i-1][0][0]-prices[i])
		dp[i][2][0] = max(dp[i-1][2][0], dp[i-1][2][1]+prices[i])
		dp[i][2][1] = max(dp[i-1][2][1], dp[i-1][1][0]-prices[i])
	}
	return max(dp[len(dp)-1][0][0], dp[len(dp)-1][1][0], dp[len(dp)-1][2][0])
}
