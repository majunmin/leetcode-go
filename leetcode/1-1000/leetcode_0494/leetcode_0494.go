package leetcode_0494

//https://leetcode.cn/problems/target-sum/
func findTargetSumWays(nums []int, target int) int {
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
