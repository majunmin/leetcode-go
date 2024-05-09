package leetcode_0198

// https://leetcode.cn/problems/house-robber/
func rob(nums []int) int {
	return dpSolution3(nums)

}

// 根据 solution2  可以
// dp[i]  依赖于 dp[i-1] & dp[i-1] , 这部分空间还可以继续优化
// https://leetcode.cn/problems/house-robber/solution/dong-tai-gui-hua-jie-ti-si-bu-zou-xiang-jie-cjavap/
func dpSolution3(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}

	pre, cur := nums[0], maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		temp := cur
		cur = maxInt(pre+nums[i], cur)
		pre = temp
	}
	return cur
}

// 优化空间的解法
func dpSolution2(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		return nums[0]
	}
	//if len(nums) == 2 {
	//	return maxInt(nums[0], nums[1])
	//}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = maxInt(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = maxInt(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[len(dp)-1]
}

func dpSolution(nums []int) int {
	// param check
	n := len(nums)
	if n == 0 {
		return 0
	}
	steal := make([]int, n)
	noSteal := make([]int, n)
	// init state
	steal[0] = nums[0]
	noSteal[0] = 0

	// 状态转移方程
	// f1[0] = nums[0], f2[0] = 0
	// f1[i] = f2[i-1] + nums[i]
	// f2[i] = max(f1[i-1], f2[i-1])
	for i := 1; i < n; i++ {
		steal[i] = noSteal[i-1] + nums[i]
		noSteal[i] = maxInt(steal[i-1], noSteal[i-1])
	}
	return maxInt(steal[n-1], noSteal[n-1])
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
