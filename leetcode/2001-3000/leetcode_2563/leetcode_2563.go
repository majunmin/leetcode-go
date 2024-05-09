package leetcode_2563

import "sort"

// https://leetcode.cn/problems/count-the-number-of-fair-pairs/
func countFairPairs(nums []int, lower int, upper int) int64 {
	// 与顺序无关
	var (
		n      = len(nums)
		result int64
	)

	sort.Ints(nums)
	for i := 0; i < n-1; i++ {
		if nums[i] > upper {
			break
		}
		l := lowerBound(nums, i+1, n-1, lower-nums[i])
		r := lowerBound(nums, i+1, n-1, upper-nums[i]+1) - 1
		result += r - l + 1
	}
	return result
}

func lowerBound(nums []int, left int, right int, target int) int64 {
	left--
	right++
	for left+1 < right {
		mid := left + (right-left)>>1
		if target <= nums[mid] {
			right = mid
		} else {
			left = mid
		}
	}
	return int64(right)
}
