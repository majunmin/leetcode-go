package leetcode_0741

import "math"

// https://leetcode.cn/problems/cherry-pickup/description/
func CherryPickup(grid [][]int) int {
	var (
		INF = math.MinInt
		n   = len(grid)
		dp  = make([][][]int, 2*n-1)
	)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, n)
			for k := range dp[i][j] {
				dp[i][j][k] = INF
			}
		}
	}
	// init state
	dp[0][0][0] = grid[0][0]
	for k := 1; k < 2*n-1; k++ {
		for x1 := max(k-n+1, 0); x1 <= min(k, n-1); x1++ {
			y1 := k - x1
			if grid[x1][y1] == -1 {
				continue
			}
			for x2 := x1; x2 <= min(k, n-1); x2++ {
				y2 := k - x2
				if grid[x2][y2] == -1 {
					continue
				}
				//1. 都向右
				res := dp[k-1][x1][x2]
				//2. (x1,y1) 向右, (x2,y2) 向下
				if x1 > 0 {
					res = max(res, dp[k-1][x1-1][x2])
				}
				//3. (x1,y1) 向下, (x2,y2) 向右
				if x2 > 0 {
					res = max(res, dp[k-1][x1][x2-1])
				}
				//4. 都向下
				if x1 > 0 && x2 > 0 {
					res = max(res, dp[k-1][x1-1][x2-1])
				}
				res += grid[x1][y1]
				if x2 != x1 { //避免重复 摘取同一个🍒
					res += grid[x2][y2]
				}
				dp[k][x1][x2] = res
			}
		}
	}
	return max(dp[2*n-2][n-1][n-1], 0)
}
