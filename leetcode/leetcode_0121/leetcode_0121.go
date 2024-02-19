package leetcode_0121

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/#/description
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minVal := math.MinInt
	mProfit := 0
	for i := 0; i < len(prices); i++ {
		mProfit = max(mProfit, prices[i]+minVal)
		minVal = max(minVal, -prices[i])
	}
	return mProfit
}

func maxProfit2(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	minVal := prices[0]
	var mProfit int
	for i := 1; i < len(prices); i++ {
		mProfit = max(mProfit, prices[i]-minVal)
		minVal = max(minVal, prices[i])
	}
	return mProfit
}
