package leetcode_2401

// https://leetcode.cn/problems/longest-nice-subarray/
func longestNiceSubarray(nums []int) int {
	// param check
	if len(nums) < 2 {
		return len(nums)
	}
	var (
		mask   int
		l, r   int
		n      = len(nums)
		result int
	)
	for r < n {
		for mask&nums[r] > 0 {
			mask -= nums[l]
			l++
		}
		mask += nums[r]
		result = max(result, r-l+1)
		r++
	}
	return result
}
