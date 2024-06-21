package leetcode_2316

// https://leetcode.cn/problems/count-unreachable-pairs-of-nodes-in-an-undirected-graph/
func countPairs(n int, edges [][]int) int64 {
	var (
		adj        = make([][]int, n)
		visited    = make([]bool, n)
		countNodes func(node int) int
		total      int
		result     int
	)

	countNodes = func(node int) int {
		//
		if visited[node] {
			return 0
		}
		visited[node] = true
		result := 1
		for _, next := range adj[node] {
			result += countNodes(next)
		}
		return result
	}

	//构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		cnt := countNodes(i)
		result += cnt * total
		total += cnt
	}
	return int64(result)
}
