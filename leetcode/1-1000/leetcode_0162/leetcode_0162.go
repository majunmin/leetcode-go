package leetcode_0162

// https://leetcode.cn/problems/find-peak-element/?envType=study-plan-v2&envId=top-interview-150
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)>>1
		if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[right]
}
