package leetcode_1984

import (
	"math"
	"sort"
)

// https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/
func minimumDifference(nums []int, k int) int {
	diff := math.MaxInt
	sort.Ints(nums)
	l, r := 0, k-1
	for l < len(nums)-k+1 {
		diff = min(diff, nums[r]-nums[l])
	}
	return diff
}
