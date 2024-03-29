package leetcode_0034

func searchRange(nums []int, target int) []int {
	// param check
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := searchLeft(nums, target)
	if nums[left] != target {
		return []int{-1, -1}
	}
	right := searchRight(nums, target)
	return []int{left, right}
}

func searchLeft(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)>>1
		if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

func searchRight(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l+1)>>1
		if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid
		}
	}
	return l
}
