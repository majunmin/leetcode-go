package leetcode_1438

// https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
func longestSubarray(nums []int, limit int) int {
	// param check
	if len(nums) < 2 {
		return len(nums)
	}
	var (
		n = len(nums)
		// 使用单调栈维护窗口内的 最大值和最小值
		maxStack, minStack = make([]int, 0, n), make([]int, 0, n)
		l, r               = 0, 0
		result             int
	)
	for r < n {
		for len(minStack) > 0 && minStack[len(minStack)-1] > nums[r] {
			// pop
			minStack = minStack[:len(minStack)-1]
		}
		for len(maxStack) > 0 && maxStack[len(maxStack)-1] < nums[r] {
			// pop
			maxStack = maxStack[:len(maxStack)-1]
		}
		minStack = append(minStack, nums[r])
		maxStack = append(maxStack, nums[r])
		// shrink window if not invalid
		for len(minStack) > 0 && len(maxStack) > 0 && maxStack[0]-minStack[0] > limit {
			if minStack[0] == nums[l] {
				minStack = minStack[1:]
			}
			if maxStack[0] == nums[l] {
				maxStack = maxStack[1:]
			}
			l++
		}
		result = max(result, r-l+1)
		r++
	}
	return result
}
