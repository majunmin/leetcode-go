package leetcode_0930

// https://leetcode.cn/problems/binary-subarrays-with-sum/
func numSubarraysWithSum(nums []int, goal int) int {
	var (
		preSumCnt = make(map[int]int) //记录前缀和出现的次数
		preSum    = 0
		result    = 0
	)
	preSumCnt[0] = 1
	for _, num := range nums {
		// 计算当前的preSum
		preSum += num
		// 统计结果
		result += preSumCnt[preSum-goal]
		preSumCnt[preSum] += 1
	}
	return result
}
