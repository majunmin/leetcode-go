package __bottomup

// https://leetcode.cn/problems/cherry-pickup-ii/
// 2.自顶向下的解法
func cherryPickup(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		r, c = len(grid), len(grid[0])
		dp   = make([][][]int, r+1)
	)
	for i := range dp {
		dp[i] = make([][]int, c)
		for j := range dp[i] {
			dp[i][j] = make([]int, c)
			//for k := range dp[i][j] {
			//	dp[i][j][k] = 0
			//}
		}
	}
	for y := r - 1; y >= 0; y-- {
		for x1 := 0; x1 < c; x1++ {
			for x2 := 0; x2 < c; x2++ {
				cur := grid[y][x1]
				if x1 != x2 {
					cur += grid[y][x2]
				}
				for d1 := -1; d1 <= 1; d1++ {
					for d2 := -1; d2 <= 1; d2++ {
						// check valid
						if x1+d1 < 0 || x1+d1 >= c || x2+d2 < 0 || x2+d2 >= c {
							continue
						}
						dp[y][x1][x2] = max(dp[y][x1][x2], cur+dp[y+1][x1+d1][x2+d2])
					}
				}
			}
		}
	}
	return dp[0][0][c-1]
}
