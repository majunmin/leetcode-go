package leetcode_1553

import "fmt"

// 超出内存限制
func minDays(n int) int {
	// param check
	if n <= 0 {
		return -1
	}
	var (
		dp = make([]int, n+1)
	)
	// init state
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = i
		//吃一个橘子🍊
		dp[i] = min(dp[i], dp[i-1]+1)
		// 橘子数可以被2整除
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		// 橘子数可以被3整除
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}
	return dp[n]
}

func minDays2(n int) int {
	var (
		memo = make(map[int]int)
	)
	memo[1] = 1
	var dfs func(int) int
	dfs = func(i int) int {
		// terminate
		if memo[i] > 0 {
			return memo[i]
		}
		step := dfs(i-1) + 1
		if i%2 == 0 {
			step = min(step, dfs(i/2)+1)
		}
		if i%3 == 0 {
			step = min(step, dfs(i/3)+1)
		}
		memo[i] = step
		return step
	}
	dfs(n)
	fmt.Println(memo)
	return memo[n]
}
