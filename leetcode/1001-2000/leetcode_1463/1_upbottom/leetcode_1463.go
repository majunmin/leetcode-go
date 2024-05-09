package __upbottom

// https://leetcode.cn/problems/cherry-pickup-ii/
func cherryPickup(grid [][]int) int {
	// param check
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	var (
		r, c = len(grid), len(grid[0])
		memo = make([][][]int, r)
	)
	for i := 0; i < len(memo); i++ {
		memo[i] = make([][]int, c)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = make([]int, c)
			for k := 0; k < c; k++ {
				memo[i][j][k] = -1
			}
		}
	}
	// 第 y 行, 连个机器人处于位置 x1, x2 能采集到的最多的樱桃🍒数
	var dfs func(y, x1, x2 int) int
	dfs = func(y, x1, x2 int) int {
		// terminate
		if x1 < 0 || x1 >= c || x2 < 0 || x2 >= c || y >= r {
			return 0
		}
		ans := memo[y][x1][x2]
		// 表示已经计算过了。
		if ans != -1 {
			return ans
		}
		cur := grid[y][x1]
		if x1 != x2 {
			cur += grid[y][x2]
		}

		// 3 * 3 = 9中场景
		// memo[y][x1][x2] = max(dfs(y+1,x1-1,x2-1),dfs(y+1,x1,x2-1), dfs(y+1,x1+1,x2-1),
		//                       dfs(y+1,x1-1,x2),dfs(y+1,x1,x2), dfs(y+1,x1+1,x2),
		//                       dfs(y+1,x1-1,x2+1),dfs(y+1,x1,x2+1), dfs(y+1,x1+1,x2+1),
		for d1 := -1; d1 <= 1; d1++ {
			for d2 := -1; d2 <= 1; d2++ {
				ans = max(ans, cur+dfs(y+1, x1+d1, x2+d2))
			}
		}
		// 记忆化
		memo[y][x1][x2] = ans
		return ans
	}
	return dfs(0, 0, c-1)
}
