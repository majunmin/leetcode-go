package leetcode_0416

func canPartition(nums []int) bool {
	return dpSolution2(nums)
}

// 空间优化
func dpSolution2(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	// 如果 总和奇数, 返回 false
	if total&1 == 1 {
		return false
	}
	// 背包问题, nums 数组 能不能将容量为total/2的背包装满
	total >>= 1
	dp := make([]bool, total+1)
	// init state
	dp[0] = true
	for _, num := range nums { // 遍历物品
		for i := len(dp) - 1; i >= 0; i-- {
			if i-num >= 0 {
				dp[i] = dp[i-num]
			}
		}
	}
	return dp[len(dp)-1]
}

func dpSolution1(nums []int) bool {
	var total int
	for _, num := range nums {
		total += num
	}
	// 如果 总和奇数, 返回 false
	if total&1 == 1 {
		return false
	}
	// 背包问题, nums 数组 能不能将容量为total/2的背包装满
	total >>= 1
	// 0-1 背包问题
	dp := make([][]bool, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		dp[i] = make([]bool, total+1)
	}
	dp[0][0] = true
	for i := 1; i < len(dp); i++ {
		for j := len(dp) - 1; j >= 0; j-- {
			dp[i][j] = dp[i-1][j]
			if j+nums[i] < len(dp) && dp[i][j] {
				dp[i][j+nums[i]] = dp[i][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}
