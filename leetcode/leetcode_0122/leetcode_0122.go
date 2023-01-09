package leetcode_0122

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/submissions/
func maxProfit(prices []int) int {

	return dpSolution2(prices)
}

// 在 solution1 的基础上 优化空间,
// 某一天的状态仅仅与前一天的状态有关, 所以  可以做如下优化
func dpSolution2(prices []int) int {
	// param check
	if len(prices) <= 1 {
		return 0
	}

	// init state, 第一天
	holdStock := -prices[0]
	holdCash := 0

	for i := 0; i < len(prices); i++ {
		preHoldStock := holdStock
		preHoldCash := holdCash
		holdStock = maxInt(preHoldStock, preHoldCash-prices[i])
		holdCash = maxInt(preHoldCash, preHoldStock+prices[i])
	}
	return holdCash
}

// 状态 为  持有 股票 & 非持有股票 是  持有的最大金币数
// 某天的行为
//      - 买入股票
//      - 卖出股票
//      - 什么也不做
// dp[0][0] = -prices[0]
// dp[0][1] = 0
// dp[i][0] = max(dp[i-1][0], dp[i-1][1] - prices[0])
// dp[i][0] = max(dp[i-1][1], dp[i-1][0] + prices[0])

func dpSolution1(prices []int) int {
	n := len(prices)
	// param check
	if n <= 1 {
		return 0
	}

	//
	dp := make([][2]int, n)
	// dp[i][0] 持有股票的 最大收益
	// dp[i][1] 不持有股票的最大收益
	dp[0][0], dp[0][1] = -prices[0], 0
	for i := 1; i < n; i++ {
		dp[i][0] = maxInt(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = maxInt(dp[i-1][1], dp[i-1][0]+prices[i])
	}
	return dp[n-1][1]
}

func maxInt(a, b int) int {
	if a > b {

		return a
	}
	return b
}
