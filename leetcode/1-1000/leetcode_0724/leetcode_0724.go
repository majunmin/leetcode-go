package leetcode_0724

// https://leetcode.cn/problems/find-pivot-index/
func pivotIndex(nums []int) int {

	// 前缀和
	var (
		n      = len(nums)
		preSum = make([]int, n+1)
	)
	for i := 0; i < n; i++ {
		preSum[i+1] = preSum[i] + nums[i]
	}
	for i := 0; i < n; i++ {
		if preSum[i] == preSum[n]-preSum[i] {
			return i
		}
	}
	return -1
}

// solution2
// result i 应该满足
// sum  + nums[i] + left(== num) == total
func pivotIndex2(nums []int) int {

	var (
		total int
		sum   int
	)
	for _, num := range nums {
		total += num
	}
	for i := 0; i < len(nums); i++ {
		if sum*2+nums[i] == total {
			return i
		}
		sum += nums[i]
	}
	return -1
}
