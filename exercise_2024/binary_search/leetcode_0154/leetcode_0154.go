package leetcode_0154

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array-ii/
func findMin(nums []int) int {
	n := len(nums)
	if n == 0 {
		panic("invalid param")
	}
	if n == 1 {
		return nums[0]
	}
	//
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] < nums[n-1] {
			right = mid
		} else if nums[mid] > nums[n-1] {
			left = mid + 1
		} else {
			left += 1
		}
	}
	return left
}
