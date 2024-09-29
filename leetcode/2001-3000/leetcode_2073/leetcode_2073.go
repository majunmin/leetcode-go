package leetcode_2073

// https://leetcode.cn/problems/time-needed-to-buy-tickets/description/
func timeRequiredToBuy(tickets []int, k int) int {
	var result int
	tk := tickets[k]
	for i, t := range tickets {
		if i <= k {
			result += min(t, tk)
		} else {
			result += min(t, tk-1)
		}
	}
	return result
}
