package leetcode_0121

func maxProfit(prices []int) int {
	// param check
	if len(prices) < 2 {
		return 0
	}
	var result int
	minVal := prices[0]
	for i := 1; i < len(prices); i++ {
		result = max(result, prices[i]-minVal)
		minVal = min(minVal, prices[i])
	}
	return result
}
