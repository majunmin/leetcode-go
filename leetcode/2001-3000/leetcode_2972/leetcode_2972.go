package leetcode_2972

// 2972. 统计移除递增子数组的数目 II
// https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/
func incremovableSubarrayCount(nums []int) int64 {
	var (
		n = len(nums)
		l int
	)
	for l < n-1 && nums[l] <= nums[l+1] {
		l++
	}
	// 整个数组是递增的  if l == n-1
	if l == n-1 {
		return int64(n*(n+1)) / 2
	}
	result := int64(l) + 2
	r := n - 1
	for r == n-1 || nums[r] > nums[r+1] {
		for l >= 0 && nums[l] >= nums[r] {
			l -= 1
		}
		result += int64(l) + 2
		r--
	}

	return result
}
