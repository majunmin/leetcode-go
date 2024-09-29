package leetcode_2376

import (
	"strconv"
)

// https://leetcode.cn/problems/count-special-integers/description/
func countSpecialNumbers(n int) int {
	s := strconv.Itoa(n)
	var (
		size = len(s)
		//mask int
		// 第i位.  s visited
		dfs  func(int, int, bool, bool) int
		memo = make([][1 << 10]int, size)
	)
	// init memo
	for i := 0; i < size; i++ {
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	dfs = func(i int, mask int, isLimit bool, isNum bool) int {
		// termiante
		if i == size {
			if isNum { // 找到一个合法的数字
				return 1
			}
			return 0
		}
		//
		if !isLimit && isNum && memo[i][mask] != -1 {
			return memo[i][mask]
		}
		var res int
		// 可以跳过当前数位.
		if !isNum {
			res = dfs(i+1, mask, false, false)
		}
		// 枚举 [d, up]
		d := 0
		if !isNum {
			d = 1
		}
		up := 9
		if isLimit {
			up = int(s[i] - '0')
		}
		for ; d <= up; d++ {
			if mask&(1<<d) != 0 {
				continue
			}
			res += dfs(i+1, mask|(1<<d), isLimit && d == up, true)
		}
		if !isLimit && isNum {
			memo[i][mask] = res
		}
		return res
	}
	return dfs(0, 0, true, false)
}
