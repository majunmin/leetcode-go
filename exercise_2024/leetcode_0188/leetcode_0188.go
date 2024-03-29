package leetcode_0188

import "math"

func maxProfit(k int, prices []int) int {
	// param check
	length := len(prices)
	// process
	dp := make([][][2]int, length+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([][2]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j][1] = math.MinInt
		}
	}
	// init
	for i := 1; i < len(dp); i++ {
		//dp[i][0][1] = -prices[i-1]
		for j := 1; j <= k; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i-1])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i-1])
		}
	}
	//var result int
	//for i := 0; i <= k; i++ {
	//	result = max(result, dp[len(dp)-1][i][0])
	//}
	//return result
	// 由于结果带有前缀性质，所以这里直接返回  dp[len(dp)-1][k][0]
	return dp[len(dp)-1][k][0]
}

func greedy(prices []int) int {
	var result int
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			result += (prices[i] - prices[i-1])
		}
	}
	return result
}
