package leetcode_0121

import "math"

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/#/description
func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	min := math.MinInt
	mProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
			continue
		}
		mProfit = maxInt(mProfit, prices[i]-min)
	}
	return mProfit
}

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
