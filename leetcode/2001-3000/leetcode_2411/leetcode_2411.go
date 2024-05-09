package leetcode_2411

// https://leetcode.cn/problems/smallest-subarrays-with-maximum-bitwise-or/
func smallestSubarrays(nums []int) []int {
	// param check
	if len(nums) == 0 {
		return nil
	}
	n := len(nums)
	result := make([]int, n)
	for i := 0; i < n; i++ {
		//nums[i] | nums[j] != nums[j] 说明 nums[i] | nums[j]  > nums[j]
		for j := i - 1; j >= 0 && nums[i]|nums[j] != nums[j]; j-- {
			nums[j] |= nums[i]
			result[j] = i - j + 1
		}
	}
	return result
}

func solution1(nums []int) []int {
	n := len(nums)
	dp := make([]int, 32)
	for i := range dp {
		dp[i] = -1
	}
	result := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		maxVal := 1
		for j := 0; j < 32; j++ {
			if nums[i]>>j&1 == 1 {
				dp[j] = i
			}
			if dp[j] != -1 {
				maxVal = max(maxVal, dp[j]-i+1)
			}
		}
		result[i] = maxVal
	}
	return result
}
