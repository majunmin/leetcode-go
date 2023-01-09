package leetcode_0033

// https://leetcode.cn/problems/search-in-rotated-sorted-array/
func search(nums []int, target int) int {
	// 没找到
	if len(nums) == 0 {
		return -1
	}
	//
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] <= nums[right] {
			if nums[mid] <= target && target <= nums[right] {
				left = mid
			} else {
				right = mid - 1
			}
		} else {
			if nums[left] <= target && target <= nums[mid-1] {
				right = mid - 1
			} else {
				left = mid
			}
		}
	}
	if nums[left] == target {
		return left
	}
	return -1
}
