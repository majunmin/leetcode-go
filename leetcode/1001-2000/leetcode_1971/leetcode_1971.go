package leetcode_1971

// https://leetcode.cn/problems/find-if-path-exists-in-graph/
// 1. 并查集
// 2.  dfs /bfs 图的遍历

// 1. bfs solution
func validPath(n int, edges [][]int, source int, destination int) bool {
	var (
		adj     = make([][]int, n)
		dfs     func(int, int) bool
		visited = make(map[int]bool)
	)
	//构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	// 判断两个节点是否
	dfs = func(cur int, target int) bool {
		// terminate
		if cur == target {
			return true
		}
		visited[cur] = true
		for _, next := range adj[cur] {
			if visited[next] {
				continue
			}
			if dfs(next, target) {
				return true
			}
		}
		return false
	}
	return dfs(source, destination)
}

func validPath2(n int, edges [][]int, source int, destination int) bool {
	var (
		adj     = make([][]int, n)
		queue   = make([]int, 0)
		visited = make([]bool, n)
	)
	//构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	// 判断两个节点是否
	queue = append(queue, source)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			if item == destination {
				return true
			}
			visited[item] = true
			for _, next := range adj[item] {
				if visited[next] {
					continue
				}
				queue = append(queue, next)
			}
		}
		queue = queue[size:]
	}

	return false
}

// unionFind
