package leetcode_3040

// https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/?envType=daily-question&envId=2024-06-08
func maxOperations(nums []int) int {
	var (
		n    = len(nums)
		memo = make([][]int, n)
		l, r = 0, n - 1
	)
	for i := 0; i < n; i++ {
		memo[i] = make([]int, n)
		for j := 0; j < n; j++ {
			memo[i][j] = -1
		}
	}

	var dfs func(left, right, target int) int
	dfs = func(left, right, target int) int {
		if left >= right {
			return 0
		}
		if memo[left][right] != -1 {
			return memo[left][right]
		}
		var res int
		if nums[left]+nums[right] == target {
			res = max(res, dfs(left+1, right-1, target)+1)
		}
		if nums[left]+nums[left+1] == target {
			res = max(res, dfs(left+2, right, target)+1)
		}
		if nums[right-1]+nums[right] == target {
			res = max(res, dfs(left, right-2, target)+1)
		}
		memo[left][right] = res
		return res
	}

	res1 := dfs(l+1, r-1, nums[l]+nums[r])
	res2 := dfs(l+2, r, nums[l]+nums[l+1])
	res3 := dfs(l, r-2, nums[r]+nums[r-1])
	return max(res1, res2, res3) + 1
}
