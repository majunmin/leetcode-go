package leetcode_0096

// https://leetcode.cn/problems/unique-binary-search-trees/
func numTrees(n int) int {
	if n < 2 {
		return 1
	}
	// 动态规划
	dp := make([]int, n+1)
	// init state
	dp[0] = 1
	dp[1] = 1
	// 状态转移方程
	//dp[i]  = sum(dp[j - 1] *  dp[j-k]) k(1...j)
	for i := 2; i <= n; i++ {
		for j := 1; j <= n; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}

// 带记忆的递归， 减少了重复计算
func recursiveSolution2(n int) int {
	// 动态规划, 枚举顶点, 然后累加结果
	mem := make(map[int]int)
	mem[0] = 1
	mem[1] = 1
	return numTreesHelp(n, mem)
}

func numTreesHelp(n int, mem map[int]int) int {
	//
	if res, exist := mem[n]; exist {
		return res
	}
	var cnt int
	for i := 1; i <= n; i++ {
		cnt += numTreesHelp(i-1, mem) * numTreesHelp(n-i, mem)
	}
	mem[n] = cnt
	return cnt
}

// 普通递归解法, 可能会超时, 因为有大量的重入计算,可以引入记忆化
func recursiveSolution(n int) int {
	if n == 1 || n == 0 {
		return 1
	}
	var sum int
	for i := 1; i <= n; i++ {
		sum += numTrees(i-1) * numTrees(n-i)
	}
	return sum
}
