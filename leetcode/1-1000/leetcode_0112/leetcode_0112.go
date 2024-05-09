package leetcode_0112

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/
func maxProfit(prices []int) int {
	return dpSolution2(prices)
}

/*
*
因为 当天状态 仅依赖于 前一天的状态
所以 用一维数组 记录 前一天的状态即可， 优化空间占用
*/
func dpSolution2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	preHasStore := -prices[0]
	preNoStore := 0
	for i := 1; i < len(prices); i++ {
		preHasStore = maxInt(preHasStore, preNoStore-prices[i])
		preNoStore = maxInt(preNoStore, preHasStore+prices[i])
	}
	return preNoStore
}

/*
动态 规划解法
1. 状态定义
dp[i][0] 第 i 天 不持有股票的 **最大收益**
dp[i][1] 第 i 天 持有股票的 **最大收益**

2. 状态转移

第 i 天不持有股票,   前一天就不持有股票,或者前一天 持有股票, i 天 卖出
dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i])  // i > 0
dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])  // i > 0

3. init State
dp[0][0] = 0
dp[0][1] = -prices[0]
*/
func dpSolution1(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][2]int, 0, len(prices))
	// init state
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	// 递推 方程
	for i := 1; i < len(prices); i++ {
		dp[i][0] = maxInt(dp[i][0], dp[i-1][1]+prices[i])
		dp[i][1] = maxInt(dp[i][1], dp[i-1][0]-prices[i])
	}
	return dp[len(dp)-1][0]
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func tanxinSolution(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	mProfit := 0
	for i := 1; i < len(prices); i++ {
		if p := prices[i] - prices[i-1]; p > 0 {
			mProfit += p
		}
	}

	return mProfit
}
