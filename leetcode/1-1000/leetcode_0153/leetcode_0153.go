package leetcode_0153

// https://leetcode.cn/problems/find-minimum-in-rotated-sorted-array/
func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	//if len(nums) == 1 {
	//	return nums[0]
	//}
	n := len(nums)
	left, right := 0, n-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return nums[left]
}

// 如何 通过 最大值 寻找最小值：  (left + 1)%len(nums)
func findMax(nums []int) int {
	if len(nums) == 0 {
		return -1
	}

	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)>>1
		if nums[mid] > nums[left] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}
