package _bottomup2

// https://leetcode.cn/problems/cherry-pickup-ii/
// 3.自顶向下的解法 的空间优化解法
// 由于 每一行的求解只与下一行相关, 所有空间上可以优化成二维
func cherryPickup(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		r, c = len(grid), len(grid[0])
		dp   = make([][]int, c)
	)
	for i := range dp {
		dp[i] = make([]int, c)
	}
	for y := r - 1; y >= 0; y-- {
		tmp := makeLevel(c)
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
						tmp[x1][x2] = max(tmp[x1][x2], cur+dp[x1+d1][x2+d2])
					}
				}
			}
		}
		dp = tmp
	}
	return dp[0][c-1]
}
func makeLevel(n int) [][]int {
	nums := make([][]int, n)
	for i := range nums {
		nums[i] = make([]int, n)
	}
	return nums
}
