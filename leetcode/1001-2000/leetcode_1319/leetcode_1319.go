package leetcode_1319

// https://leetcode.cn/problems/number-of-operations-to-make-network-connected/
func makeConnected(n int, connections [][]int) int {
	// param check
	if len(connections) < n-1 {
		return -1
	}
	//统计连通图的个数 cnt, 结果就是 cnt-1
	var (
		adj      = make([][]int, n)
		visited  = make([]bool, n)
		groupNum int
		dfs      func(int)
	)
	dfs = func(node int) {
		visited[node] = true
		for _, next := range adj[node] {
			if visited[next] {
				continue
			}
			dfs(next)
		}
	}
	//构建图
	for i := 0; i < len(connections); i++ {
		adj[connections[i][0]] = append(adj[connections[i][0]], connections[i][1])
		adj[connections[i][1]] = append(adj[connections[i][1]], connections[i][0])
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		groupNum++
		dfs(i)
	}
	return groupNum - 1
}
