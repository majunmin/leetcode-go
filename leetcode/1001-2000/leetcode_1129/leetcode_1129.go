package leetcode_1129

type edge struct {
	from  int
	to    int
	color int
}

// color: red 1 blue 2 unknown 0
func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	var (
		result  = make([]int, 0)
		d       int
		queue   = make([][2]int, 0)
		adj     = make([][]edge, 0)
		visited = make([][3]bool, n)
	)
	for i := 0; i < n; i++ {
		result[i] = -1
	}
	for _, e := range redEdges {
		adj[e[0]] = append(adj[e[0]], edge{from: e[0], to: e[1], color: 1})
	}
	for _, e := range blueEdges {
		adj[e[0]] = append(adj[e[0]], edge{from: e[0], to: e[1], color: 0})
	}
	queue = append(queue, [2]int{0, 1})
	queue = append(queue, [2]int{0, 0})
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			cur, c := item[0], item[1]
			if result[cur] == -1 {
				result[cur] = d
			}
			visited[cur][c] = true
			c = 1 - c
			for _, e := range adj[cur] {
				next, nextColor := e.to, e.color
				if visited[next][nextColor] {
					queue = append(queue, [2]int{next, nextColor})
				}
			}
		}
		queue = queue[size:]
		d++
	}
	return result
}

func dfsSolution(n int, redEdges [][]int, blueEdges [][]int) []int {
	var (
		result  = make([]int, n)
		adj     = make([][]edge, n)
		visited = make([][3]bool, n)
	)

	// 构建图
	for _, e := range redEdges {
		adj[e[0]] = append(adj[e[0]], edge{from: e[0], to: e[1], color: 1})
	}
	for _, e := range blueEdges {
		adj[e[0]] = append(adj[e[0]], edge{from: e[0], to: e[1], color: 2})
	}
	for i := 1; i < n; i++ {
		result[i] = -1
	}
	visited[0][0] = true
	dfs(0, adj, 0, 0, result, visited)
	return result
}

func dfs(cur int, adj [][]edge, cnt int, color int, result []int, visited [][3]bool) {
	// terminate
	for _, e := range adj[cur] {
		next, nextColor := e.to, e.color
		if nextColor == color {
			continue
		}
		if result[next] == -1 || cnt+1 < result[next] {
			result[next] = cnt + 1
		}
		if visited[next][nextColor] {
			continue
		}
		visited[next][nextColor] = true
		dfs(next, adj, cnt+1, nextColor, result, visited)
		visited[next][nextColor] = false
	}
}
