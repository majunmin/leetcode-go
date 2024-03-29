package leetcode_2908

import "math"

// https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i/
func minimumSum(nums []int) int {
	return solution1(nums)
}

// 类似于题目 接雨水
func solution1(nums []int) int {
	// param check
	if len(nums) < 3 {
		return 0
	}
	//
	leftMin, rightMin := make([]int, len(nums)), make([]int, len(nums))
	leftMin[0] = nums[0]
	rightMin[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums)-1; i++ {
		leftMin[i] = min(leftMin[i-1], nums[i])
		rightMin[len(nums)-i-1] = min(rightMin[len(nums)-i], nums[len(nums)-i-1])
	}
	minSum := math.MaxInt
	for i := 1; i < len(nums)-1; i++ {
		if nums[i] > leftMin[i] && nums[i] > rightMin[i] {
			minSum = min(minSum, nums[i]+leftMin[i]+rightMin[i])
		}
	}
	if minSum == math.MaxInt {
		return -1
	}
	return minSum
}
