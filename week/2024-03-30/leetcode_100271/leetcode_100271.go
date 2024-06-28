package leetcode_100271

import "math"

// https://leetcode.cn/contest/biweekly-contest-127/problems/shortest-subarray-with-or-at-least-k-ii/
func minimumSubarrayLength(nums []int, k int) int {
	// param check
	if len(nums) == 0 {
		return -1
	}
	//滑动窗口解法，r指针滑动很好处理， 如何处理l指针的滑动呢? 剔除 nums[l]
	// a | b | c  -- a:

	var (
		total  int
		minLen = math.MaxInt
		bits   = make([]int, 32)
		l, r   int
	)
	for r < len(nums) {
		total |= nums[r]
		for total < k {
			// todo: remove nums[l]
			l++
		}
		if total >= k {
			minLen = min(minLen, r-l+1)
		}
	}
	if minLen == math.MaxInt {
		return -1
	}
	return minLen
}
