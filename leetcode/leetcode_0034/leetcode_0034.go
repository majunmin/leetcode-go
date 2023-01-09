package leetcode_0034

// https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/
func searchRange(nums []int, target int) []int {
	// param check
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeft(nums, target)
	if nums[left] != target {
		left = -1
	}
	right := findRight(nums, target)
	if nums[right] != target {
		right = -1
	}
	return []int{left, right}
}

func findRight(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left+1)/2
		if nums[mid] <= target {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

func findLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
