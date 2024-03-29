package leetcode_0153

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	// param check
	if len(nums) == 0 {
		panic("invalid param")
	}
	if len(nums) == 1 {
		return nums[0]
	}
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] <= nums[len(nums)-1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
