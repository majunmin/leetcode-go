package leetcode_0600

import (
	"math/bits"
)

func findIntegers(n int) int {
	var (
		m    = bits.Len(uint(n))
		memo = make([][2]int, m)
		dfs  func(i, pre int, isLimit bool) int
	)
	// init
	for i := 0; i < m; i++ {
		memo[i] = [2]int{-1, -1}
	}
	dfs = func(i, pre int, isLimit bool) int {
		// terminate
		if i == m {
			return 1
		}
		if !isLimit && memo[i][pre] != -1 {
			return memo[i][pre]
		}
		up := 1
		if isLimit {
			up = n >> i & 1
		}
		// 填 0
		res := dfs(i+1, 0, isLimit && up == 0)
		if pre == 0 && up == 1 {
			//可以填1
			res += dfs(i+1, 1, isLimit)
		}
		return res
	}
	return dfs(0, 0, true)
}
