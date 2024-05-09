package leetcode_0154

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// 有边界 缩小 1
			right--
		}
	}
	return left
}
