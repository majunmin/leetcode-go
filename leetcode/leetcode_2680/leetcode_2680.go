package leetcode_2680

// https://leetcode.cn/problems/maximum-or/
func maximumOr(nums []int, k int) int64 {
	// param check
	if len(nums) < 0 {
		return -1
	}
	n := len(nums)
	preOr := make([]int, n)
	suffOr := make([]int, n)
	preOr[0] = 0
	suffOr[n-1] = 0
	for i := 0; i < n; i++ {
		if i > 0 {
			preOr[i] = preOr[i-1] | nums[i-1]
		}
		if i < n-1 {
			suffOr[n-i-2] = suffOr[n-i-1] | nums[n-i-1]
		}
	}
	var maxVal int
	for i := 0; i < n; i++ {
		maxVal = max(maxVal, preOr[i]|suffOr[i]|nums[i]<<k)
	}
	return int64(maxVal)
}
