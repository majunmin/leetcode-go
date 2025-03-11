package leetcode_2389

import "slices"

// https://leetcode.cn/problems/longest-subsequence-with-limited-sum/
func answerQueries(nums []int, queries []int) []int {
	slices.SortFunc(nums, func(a, b int) int {
		return a - b
	})
	var (
		result = make([]int, 0, len(queries))
		preSum = make([]int, len(nums)+1)
	)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	for _, num := range queries {
		result = append(result, upperBound(preSum, num))
	}
	return result
}

func upperBound(nums []int, target int) int {
	l, r := 0, len(nums)
	for l+1 < r {
		mid := l + (r-l)/2
		if nums[mid] > target {
			r = mid
		} else {
			l = mid
		}
	}
	return l
}
