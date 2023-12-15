package week100114

import "math"

func minimumSum(nums []int) int {
	// param check
	length := len(nums)
	if length < 3 {
		return -1
	}
	minLeft, minRight := make([]int, length), make([]int, length)
	// init
	minLeft[0] = nums[0]
	minRight[length-1] = nums[length-1]
	for i := 1; i < length; i++ {
		if nums[i] < minLeft[i-1] {
			minLeft[i] = nums[i]
		} else {
			minLeft[i] = minLeft[i-1]
		}
	}
	for i := length - 1; i >= 0; i-- {
		if nums[i] < minRight[i+1] {
			minRight[i] = nums[i]
		} else {
			minRight[i] = minRight[i+1]
		}
	}
	// find result
	result := math.MaxInt
	for i := 1; i < length-1; i++ {
		if nums[i] > minLeft[i] && nums[i] > minRight[i] {
			result = min(result, nums[i]+minLeft[i]+minRight[i])
		}
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}
