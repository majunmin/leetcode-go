package leetcode_100277

import (
	"sort"
)

func minOperationsToMakeMedianK(nums []int, k int) int64 {
	sort.Ints(nums)
	l, r := 0, len(nums)-1
	midIdx := len(nums) / 2
	result := abs(nums[midIdx] - k)
	nums[midIdx] = k
	for l < midIdx {
		if k < nums[l] {
			result += nums[l] - k
		}
		l++
	}
	for r > midIdx {
		if k > nums[r] {
			result += k - nums[r]
		}
		r--
	}
	return int64(result)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
