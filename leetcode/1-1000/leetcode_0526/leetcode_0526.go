package leetcode_0526

import "math/bits"

// https://leetcode.cn/problems/beautiful-arrangement/
func countArrangement(n int) int {
	var (
		result  int
		visited = make([]bool, n+1)
		dfs     func(int, []int)
	)
	dfs = func(idx int, path []int) {
		if len(path) == n {
			result++
			return
		}
		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			path = append(path, i)
			if idx%i == 0 || i%idx == 0 {
				dfs(idx+1, path)
			}
			path = path[:len(path)-1]
			visited[i] = false
		}
	}
	// index 从1开始编号
	dfs(1, []int{})
	return result
}

// 记忆化搜索.
// 使用 memory 记录选择n个数的 最大方案数
func countArrangement2(n int) int {
	var (
		mask = 1<<n - 1
		dfs  func(int) int
		memo = make([]int, 1<<n)
	)
	for i := 0; i < 1<<n; i++ {
		// 表示该位置还未被计算过
		memo[i] = -1
	}
	dfs = func(state int) int {
		if state == mask {
			return 1
		}
		if memo[state] != -1 {
			return memo[state]
		}
		res := 0
		idx := bits.OnesCount(uint(state)) + 1
		for j := 1; j <= n; j++ {
			if (state>>(j-1))&1 == 1 {
				continue
			}
			if j%idx != 0 && idx%j != 0 {
				continue
			}
			res += dfs(state | (1 << (j - 1)))
		}
		memo[state] = res
		return res
	}
	return memo[mask]
}

// 动态规划
// 1. 用一个整数 state(位运算状态压缩)表示选了那些数字 和未选哪些数字.
// (000101)2 表示 1 3 被使用了.
// 2. 状态定义定义
// dp[i][state] 表示 前i个数 选择了 state 的方案数.
func countArrangement3(n int) int {
	var (
		mask = 1 << n
		dp   = make([][]int, n)
	)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, mask)
	}
	//dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for state := 0; state < mask; state++ {
			// 枚举最后一位选的位置是 k
			for k := 1; k <= n; k++ {
				//该位置还没被选到过.
				if state&(1<<(k-1)) == 0 {
					continue
				}
				if k%i != 0 && i%k != 0 {
					continue
				}
				dp[i][state] += dp[i-1][state&(1<<(k-1)-1)]
			}
		}
	}
	return dp[n][mask-1]
}
