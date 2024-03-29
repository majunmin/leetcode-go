package leetcode_0560

// https://leetcode.cn/problems/subarray-sum-equals-k/?envType=study-plan-v2&envId=top-100-liked
func subarraySum(nums []int, k int) int {
	// 构建前缀和数组
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}
	//
	var result int
	for i := 1; i <= len(nums); i++ {
		for j := 0; j < i; j++ {
			if preSum[i]-preSum[j] == k {
				result++
			}
		}
	}
	return result
}

// 利用  map 优化 方法一
func solution2(nums []int, k int) int {
	prevSum := make(map[int]int) // 前缀和以及前缀和对应的 出现的次数
	prevSum[0] = 1
	var sum int
	var result int
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		result += prevSum[k-nums[i]]
		prevSum[sum]++
	}
	return result
}
