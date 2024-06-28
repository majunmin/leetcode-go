package _024_06_23

import "sort"

func minimumAverage(nums []int) float64 {
	var (
		result float64 = 50
		l, r           = 0, len(nums) - 1
	)
	sort.Ints(nums)
	for l < r {
		result = min(result, float64(nums[l]+nums[r])/2.0)
	}
	return result
}
