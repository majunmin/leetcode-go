package leetcode_3007

// https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/
func findMaximumNumber(k int64, x int) int64 {
	var (
		totalValue int
		num        int64 = 1
	)
	for num < k {
		for temp := num; temp > 0; temp = temp >> x {
			if temp>>x&1 == 1 {
				totalValue++
			}
		}
	}
	return num
}
