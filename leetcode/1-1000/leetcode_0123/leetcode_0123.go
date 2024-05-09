package leetcode_0123

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
func maxProfit(prices []int) int {
	return dpSolution(prices)
}

func dpSolution(prices []int) int {
	/// param check
	if len(prices) <= 1 {
		return 0
	}
	// 定义四个状态
	// holdStock1 第一次 持有股票的最大收益
	// holdCrash1 第一次 不持有股票的最大收益
	// holdStock2 第二次 不持有股票的最大收益 与第一次 卖出有关系
	// holdCrash2 第二次 不持有股票的最大收益

	// initState
	holdStock1, holdCrash1 := -prices[0], 0
	holdStock2, holdCrash2 := math.MinInt, math.MinInt
	for i := 1; i < len(prices); i++ {
		holdStock1 = maxInt(holdStock1, -prices[i])
		holdCrash1 = maxInt(holdCrash1, holdStock1+prices[i])
		holdStock2 = maxInt(holdStock2, holdCrash1-prices[i])
		holdCrash2 = maxInt(holdCrash2, holdStock2+prices[i])
	}
	return maxInt(holdCrash1, holdCrash2)

}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
