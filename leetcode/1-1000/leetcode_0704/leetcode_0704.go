package leetcode_0704

func search(nums []int, target int) int {
	// param check
	if len(nums) == 0 {
		return -1
	}
	var (
		left, right = -1, len(nums)
	)
	for left+1 < right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid
		} else {
			right = mid
		}
	}
	if right == len(nums) || nums[right] != target {
		return -1
	}
	return right
}
