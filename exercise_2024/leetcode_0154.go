package exercise_2024

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
func findMin2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// param check
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1 // 防止溢出
		if nums[mid] <= nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}
