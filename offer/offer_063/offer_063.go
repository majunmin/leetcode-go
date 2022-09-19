package offer_063

//https://leetcode.cn/problems/gu-piao-de-zui-da-li-run-lcof/?envType=study-plan&id=lcof
func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	maxProf := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		maxProf = maxInt(maxProf, prices[i]-minPrice)
		minPrice = minInt(prices[i], minPrice)
	}

	return maxProf
}

func minInt(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
