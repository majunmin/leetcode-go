package leetcode_3283

import (
	"math"
	"math/bits"
)

var (
	directions = [][]int{
		{-1, -2},
		{1, -2},
		{2, -1},
		{2, 1},
		{1, 2},
		{-1, 2},
		{-2, 1},
		{-2, -1},
	}
)

// https://leetcode.cn/problems/maximum-number-of-moves-to-kill-all-pawns/description/
func maxMoves(kx int, ky int, positions [][]int) int {
	// 预计算距离的时间复杂度
	// 预计算兵到每个位置的距离: 15*50*50
	// 或者与预计算马与兵的距离: 50*50*50*50  (马的位置的可能性 = 50*50)
	var (
		n   = len(positions)
		dis = make([][][]int, n)
	)
	// 1. 计算距离
	for i := 0; i < n; i++ {
		dis[i] = make([][]int, 50)
		// bfs
		bfsDistance(positions[i], dis[i])
	}
	// 2. 状态压缩dp
	// memo[i][j]: 马在第i个兵的位置时
	memo := make([][]int, n+2)
	// init memo[i][j] 表示该节点还未被访问过.
	for i := 0; i < n+2; i++ {
		memo[i] = make([]int, 1<<n)
		for j := 0; j < len(memo[i]); j++ {
			memo[i][j] = -1
		}
	}
	return dfs(n+1, kx, ky, 1<<n-1, dis, positions, memo)
}

type point struct {
	x, y int
}

// 计算兵的位置 posotion 到棋盘的每个位置的最短距离.
// - 这里其实 visited 可以省略, 用 distance[i][j] = -1 表示该位置还未被计算过.
func bfsDistance(position []int, dis [][]int) {
	var (
		queue   = make([]point, 0)
		visited = make([][]bool, 50)
		step    = 0
	)
	for i := 0; i < 50; i++ {
		dis[i] = make([]int, 50)
		visited[i] = make([]bool, 50)
	}
	x, y := position[0], position[1]
	queue = append(queue, point{x, y})
	visited[x][y] = true
	for len(queue) > 0 {
		size := len(queue)
		step++
		for i := 0; i < size; i++ {
			p := queue[i]
			for _, dir := range directions {
				newx, newy := p.x+dir[0], p.y+dir[1]
				if newx < 0 || newx >= 50 || newy < 0 || newy >= 50 || visited[newx][newy] {
					continue
				}
				dis[newx][newy] = step
				visited[newx][newy] = true
				queue = append(queue, point{x: newx, y: newy})
			}
		}
		queue = queue[size:]
	}
}

// cur = n+1表示一开始🐴的位置
// kx,ky 一开始马的位置.
// position, cur : 可以得出当前🐴的位置.
func dfs(cur, kx, ky int, mask int, dis [][][]int, positions [][]int, memo [][]int) int {
	if mask == 0 {
		return 0
	}
	// memo[cur][mask] 之前计算过.
	if memo[cur][mask] != -1 {
		return memo[cur][mask]
	}
	var (
		n   = len(positions)
		res = 0
		// 当前🐴所在位置.
		x = kx
		y = ky
	)
	if cur < n {
		x = positions[cur][0]
		y = positions[cur][1]
	}

	if (n-bits.OnesCount(uint(mask)))&1 == 1 {
		// 已经吃掉的兵的个数是奇数 -> 该bob走了
		res = math.MaxInt
		for i := 0; i < n; i++ {
			// 如果位置为i的兵还未被访问, 则 吃掉该兵,并将该兵的位置置为0
			if (mask>>i)&1 == 1 {
				res = min(res, dis[i][x][y]+dfs(i, kx, ky, mask^(1<<i), dis, positions, memo))
			}
		}
	} else {
		// 已经吃掉的兵的个数是偶数 -> 该 alice 走了
		for i := 0; i < n; i++ {
			// 如果位置为i的兵还未被访问, 则 吃掉该兵,并将该兵的位置置为0
			if mask>>i&1 == 1 {
				// 将 第i位置0
				res = max(res, dis[i][x][y]+dfs(i, kx, ky, mask^(1<<i), dis, positions, memo))
			}
		}
	}
	memo[cur][mask] = res
	return res
}
