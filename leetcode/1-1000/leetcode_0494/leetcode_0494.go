package leetcode_0494

// https://leetcode.cn/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
	return dpSolution(nums, target)
}

func dpSolution(nums []int, target int) int {
	var (
		totalSum int
	)
	for _, num := range nums {
		totalSum += num
	}
	// -totalSum ...0...totalSum
	// item = 0 ->
	//
	var (
		n  = len(nums)
		dp = make([][]int, n)
	)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2*totalSum+1)
	}
	// init state
	dp[0][nums[0]+totalSum] = 1
	dp[0][-nums[0]+totalSum] = 1
	for i := 1; i < n; i++ {
		for j := 0; j < 2*totalSum+1; j++ {
			//dp[i][j] = dp[i-1][j]
			if j-nums[i] >= 0 {
				dp[i][j] += dp[i-1][j-nums[i]]
			}
			if j+nums[i] < 2*totalSum+1 {
				dp[i][j] += dp[i-1][j+nums[i]]
			}
		}
	}
	return dp[n-1][target+totalSum]
}

// 动态规划, 背包问题🎒

func dfsSolution(nums []int, target int) int {
	var (
		result int
		//
		dfs func(int, int)
	)
	dfs = func(idx, sum int) {
		// terminate
		if idx == len(nums) {
			if sum == target {
				result++
			}
			return
		}
		dfs(idx+1, sum+nums[idx])
		dfs(idx+1, sum-nums[idx])
	}
	dfs(0, 0)
	return result
}
