package solution1

import "math"

func cherryPickup(grid [][]int) int {
	var (
		n  = len(grid)
		dp = make([][][]int, 2*n-1)
	)
	// 1. 状态定义
	// 将题目转化为 两个人同时走, 同一时刻 A处于位置(x1, y1)  B出于未知(x2, y2)
	// steps == x1 + y1 == x2 + y2
	// dp[k][x1][x2]:走了 k 步, A处于位置(x1, k-x1) B处于位置(x2, k-x2) 可以获取的樱桃🍒数的最大值
	// 罗列出所有的可能性.
	// 2. 递推公式:
	// dp[k][x1][x2] = max((x1向右, x2 向下), (x1向下, x2向下), (x1向右, x2向右), (x1向右, x2向下))
	// dp[k][x1][x2] = max(dp[k-1][x1-1][x2], dp[k-1][x1-1][x2-1], dp[k-1][x1][x2], dp[k-1][x1][x2-1])

	// ... 一些细节:

	// init state
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = math.MinInt
			}
		}
	}
	dp[0][0][0] = grid[0][0]
	//
	for k := 1; k < 2*n-1; k++ {
		for x1 := max(k-n+1, 0); x1 <= min(k, n-1); x1++ {
			y1 := k - x1
			// skip
			if grid[x1][y1] == -1 {
				continue
			}
			// ? 为什么是  x2 的初始值是 x1
			// 代码实现时，我们可以将 A 和 B 走出的路径的上轮廓看成是 A 走出的路径,下轮廓看成是 B 走出的路径,即视作 A 始终不会走到 B 的下方.
			//  则有 x1≤x2, 在代码实现时保证这一点，可以减少循环次数.
			for x2 := x1; x2 <= min(k, n-1); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				//3. (x1向右, x2向右),
				res := dp[k-1][x1][x2]
				//1. (x1向右, x2 向下),
				if x2 > 0 {
					res = max(res, dp[k-1][x1][x2-1])
				}
				//2. (x1向下, x2向下),
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[k-1][x1-1][x2-1])
				}
				//4. (x1向右, x2向下)
				if x1 > 0 {
					res = max(res, dp[k-1][x1][x2-1])
				}
				res += grid[x1][y1]
				if x1 != x2 { // 避免重复
					res += grid[x2][y2]
				}
				dp[k][x1][x2] = res
			}
		}
	}
	return min(dp[2*n-2][n-1][n-1], 0)
}
