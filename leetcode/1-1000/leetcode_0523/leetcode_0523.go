package leetcode_0523

// https://leetcode.cn/problems/continuous-subarray-sum/description/
func checkSubarraySum(nums []int, k int) bool {
	var (
		n = len(nums)
		// 前缀和数组.
		preSum    = make([]int, n+1)
		preSumSet = make(map[int]bool)
	)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	for i := 2; i <= n; i++ {
		preSumSet[preSum[i-2]%k] = true
		// 判断存在同余的数
		if preSumSet[preSum[i]%k] {
			return true
		}
	}
	return false
}
