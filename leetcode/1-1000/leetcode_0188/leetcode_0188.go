package leetcode_0188

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
func maxProfit(k int, prices []int) int {

	return dpSolution3(k, prices)
}

func dpSolution3(k int, prices []int) int {

	n := len(prices)
	if k <= 0 || n <= 1 {
		return 0
	}
	if n/2 < k {
		k = n / 2
	}

	dp := make([][2]int, k)

	for i := 0; i < k; i++ {
		dp[i][0] = -prices[i]
	}

	for i := 1; i < n; i++ {
		dp[0][0] = maxInt(dp[0][0], -prices[i])
		dp[0][1] = maxInt(dp[0][1], dp[0][0]+prices[i])

		for j := 1; j <= k; j++ {
			dp[j][0] = maxInt(dp[j][0], dp[j-1][1]-prices[i])
			dp[j][1] = maxInt(dp[j][1], dp[j][0]+prices[i])
		}
	}
	return dp[k-1][1]
}

// 优化空间的解法
func dpSolution2(k int, prices []int) int {
	n := len(prices)
	if k <= 0 || n <= 1 {
		return 0
	}

	if n/2 < k {
		k = n / 2
	}

	//
	dp := make([][2]int, k+1)
	for i := 0; i <= k; i++ {
		dp[i][0] = math.MinInt
	}

	for _, price := range prices {
		for i := 1; i <= k; i++ {
			dp[i][0] = maxInt(dp[i][0], dp[i-1][1]-price)
			dp[i][1] = maxInt(dp[i][1], dp[i][0]+price)
		}
	}

	return dp[k][1]
}

func dpSolution(k int, prices []int) int {
	n := len(prices)
	// param check
	if n <= 1 || k <= 0 {
		return 0
	}

	// 注意点之一:  n 天 内 最多可以进行 n/2 笔交易
	if n/2 < k {
		k = n / 2
	}

	// 定义状态
	// dp[i][j][0] : 第 i 天 最大 j 次交易,   持有股票 的最大收益
	// dp[i][j][1] : 第 i 天 最大 j 次交易, 不持有股票 的最大收益

	dp := make([][][2]int, n+1)
	for i := 0; i <= n; i++ {
		for j := 0; j <= k; j++ {
			dp[i] = append(dp[i], [2]int{})
			dp[i][j][0] = math.MinInt
		}
	}
	// 状态转移
	for i := 1; i <= n; i++ {
		for j := 1; j <= k; j++ {
			// max(前一天就持有股票,前一天不持有股票, 今天买入)
			dp[i][j][0] = maxInt(dp[i-1][j][0], dp[i-1][j-1][1]-prices[i-1])
			dp[i][j][1] = maxInt(dp[i-1][j][1], dp[i-1][j][0]+prices[i-1])
		}
	}
	return dp[n][k][1]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
