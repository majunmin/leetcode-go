package leetcode_2529

// https://leetcode.cn/problems/maximum-count-of-positive-integer-and-negative-integer/
func maximumCount(nums []int) int {
	l := lowerBound(nums, 0)
	r := lowerBound(nums, 1)
	return max(l, len(nums)-r)
}

func lowerBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r-l)>>1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
