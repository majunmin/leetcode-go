package leetcode_1186

// https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
func maximumSum(arr []int) int {
	// 1. 找出子数组的最大和.
	// 2. 尝试从中删除元素.(枚举遍历的方式)
	if len(arr) == 0 {
		panic("invalid param: arr.length is 0")
	}
	//dp[i][j]: 以 arr[i] 结尾 删除k次的数组的最大和.
	// init state:
	// dp[0][0] = arr[0]
	// dp[0][1] = 0
	var (
		dp     = make([][2]int, len(arr))
		result = arr[0]
	)
	//
	dp[0][0] = arr[0]
	dp[0][1] = 0
	for i := 1; i < len(arr); i++ {
		dp[i][0] = max(dp[i-1][0], 0) + arr[i]
		dp[i][1] = max(dp[i-1][1]+arr[i], dp[i-1][0])
		result = max(result, dp[i][0], dp[i][1])
	}
	return result
}
