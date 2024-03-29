package leetcode_0035

// https://leetcode.cn/problems/search-insert-position/?envType=study-plan-v2&envId=top-100-liked
func searchInsert(nums []int, target int) int {
	// 二分查找,寻找右边第一个 <= target 的值
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left+1)>>1
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid
		}
	}
	// left == right
	if left == right || target < nums[left] {
		return left
	}
	// 此时  插入位置 = len(nums)
	return left + 1
}
