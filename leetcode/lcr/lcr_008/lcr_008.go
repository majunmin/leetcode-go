package lcr_008

import "math"

// https://leetcode.cn/problems/2VG8Kg/
// 1. 前缀和 + 二分
// 2. 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	var (
		n = len(nums)
		// 前缀和数组
		preSum = make([]int, n+1)
		result = math.MaxInt
	)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
		// 找到以当前数字为结尾 && 满足条件的子数组的最小长度
		for j := i; j >= 0; j-- {
			if preSum[i+1]-preSum[j] >= target {
				result = min(result, i+1-j)
			}
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}

// 前缀和 + 二分
func minSubArrayLen2(target int, nums []int) int {
	var (
		n      = len(nums)
		preSum = make([]int, n+1)
		result = math.MaxInt
	)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	// 不存在满足条件的子数组,提前返回
	if preSum[n] < target {
		return 0
	}
	for j := 1; j <= n; j++ {
		if preSum[j]-target < 0 {
			continue
		}
		idx := upperBound(preSum, preSum[j]-target)
		result = min(result, j-idx)
	}
	return result
}

func upperBound(sum []int, target int) int {
	l, r := -1, len(sum)
	for l+1 < r {
		mid := l + (r-l)>>1
		if sum[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	return l
}

// 滑动窗口
func minSubArrayLen3(target int, nums []int) int {
	var (
		n         = len(nums)
		l, r      int
		result    = math.MaxInt
		windowSum int
	)

	for ; r < n; r++ {
		windowSum += nums[r]
		for ; l < r && windowSum >= target; l++ {
			result = min(result, r-l+1)
			windowSum -= nums[l]
		}
	}
	if result == math.MaxInt {
		return 0
	}
	return result
}
