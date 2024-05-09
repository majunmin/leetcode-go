package leetcode_1838

import "sort"

// https://leetcode.cn/problems/frequency-of-the-most-frequent-element/
func maxFrequency(nums []int, k int) int {
	sort.Ints(nums)
	var (
		result int
		l      int
		r      = 1
		total  int
	)
	for r < len(nums) {
		total += (nums[r] - nums[r-1]) * (r - l)
		// shrink window if not valid
		for total > k {
			total -= nums[r] - nums[l]
			l++
		}
		result = r - l + 1
	}
	return result
}
