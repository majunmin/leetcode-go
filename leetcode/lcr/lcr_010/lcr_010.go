package lcr_010

// https://leetcode.cn/problems/QTMn0o/
func subarraySum(nums []int, k int) int {
	var (
		preSumCache = make(map[int]int)
		preSum      int
		result      int
	)
	preSumCache[0] = 1
	for _, num := range nums {
		preSum += num
		result += preSumCache[preSum-k]
		preSumCache[preSum]++
	}

	return result
}
