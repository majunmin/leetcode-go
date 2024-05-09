package leetcode_2909

import "math"

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/
func minimumSum(nums []int) int {
	// 和 2908 题目有什么不同吗？ 没看出来
	// param check
	if len(nums) < 3 {
		return -1
	}
	n := len(nums)
	rightMin := make([]int, n)
	rightMin[n-1] = nums[n-1]
	for i := n - 2; i > 0; i-- {
		rightMin[i] = min(rightMin[i+1], nums[i])
	}
	minPre := nums[0]
	minSum := math.MaxInt
	for i := 1; i < n-1; i++ {
		minPre = min(minPre, nums[i])
		if nums[i] > minPre && nums[i] > rightMin[i] {
			minSum = min(minSum, nums[i]+minPre+rightMin[i])
		}
	}
	if minSum == math.MaxInt {
		return -1
	}
	return minSum
}
