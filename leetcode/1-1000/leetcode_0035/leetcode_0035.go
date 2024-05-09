package leetcode_0035

// https://leetcode.cn/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	// 找到第一个  >= target 的数的位置.
	if len(nums) == 0 {
		return 0
	}
	var (
		left, right = -1, len(nums)
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		//
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	return right
}
