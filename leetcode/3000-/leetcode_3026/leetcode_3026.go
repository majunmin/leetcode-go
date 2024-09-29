package leetcode_3026

import "math"

// https://leetcode.cn/problems/maximum-good-subarray-sum/
func maximumSubarraySum(nums []int, k int) int64 {
	// 前缀和 + hashMap
	// 用hashmap存储 nums[i] 和 min(s[i])的映射关系
	// 最大子数组和res 就是 max(res, s[i] - s[j-1])
	var (
		minS   = make(map[int]int)
		preSum = 0
		result = math.MinInt
	)

	for i := 0; i < len(nums); i++ {
		if _, exist := minS[nums[i]]; !exist {
			minS[nums[i]] = math.MinInt
		}
		minS[nums[i]] = min(minS[nums[i]], preSum)
		preSum += nums[i]
		// 2. 更新result
		if _, exist := minS[nums[i]-k]; exist {
			result = max(result, preSum-minS[nums[i]-k])
		}
		if _, exist := minS[nums[i]+k]; exist {
			result = max(result, preSum-minS[nums[i]+k])
		}
	}
	if result == math.MinInt {
		return 0
	}
	return int64(result)
}
