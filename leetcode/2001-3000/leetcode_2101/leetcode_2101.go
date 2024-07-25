package leetcode_2101

func maximumDetonation(bombs [][]int) int {
	if len(bombs) <= 1 {
		return len(bombs)
	}

	var (
		result = 1
		n      = len(bombs)
		// 记录 bomb之间的引爆关系
		adj = make([][]int, n)
	)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isConnected(bombs, i, j) {
				adj[i] = append(adj[i], j)
			}
		}
	}
	// 广度优先搜索
	for i := 0; i < n; i++ {
		var (
			visited = make([]bool, n)
			cnt     int
			queue   = make([]int, 0)
		)

		queue = append(queue, i)
		for len(queue) > 0 {
			size := len(queue)
			for j := 0; j < size; j++ {
				item := queue[j]
				cnt++
				for _, next := range adj[item] {
					if visited[next] {
						continue
					}
					queue = append(queue, next)
				}
			}
			queue = queue[size:]
		}
		result = max(result, cnt)
	}
	return result
}

func isConnected(bombs [][]int, from, to int) bool {
	dx := bombs[from][0] - bombs[to][0]
	dy := bombs[from][1] - bombs[to][1]
	return bombs[from][2]*bombs[from][2] >= dx*dx+dy*dy
}
