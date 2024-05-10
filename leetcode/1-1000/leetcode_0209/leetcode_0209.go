package leetcode_0209

import "math"

// https://leetcode.cn/problems/minimum-size-subarray-sum/?envType=study-plan-v2&envId=top-interview-150
func minSubArrayLen(target int, nums []int) int {
	var (
		total       int // 窗口内 数字的和
		left, right int
		result      = math.MaxInt
	)
	for ; right < len(nums); right++ {
		total += nums[right]
		// shrink window
		for left <= right && total >= target {
			result = min(result, right-left+1)
			total -= nums[left]
			left++
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
