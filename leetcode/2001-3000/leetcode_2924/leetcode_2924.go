package leetcode_2924

// https://leetcode.cn/problems/find-champion-ii/?envType=daily-question&envId=2024-04-13
func findChampion(n int, edges [][]int) int {
	// 有题目观察得知, 获得 champion 的节点的入度 == 0
	// 1. 如果 有多个节点入度是 0, 返回-1
	// 拓扑排序
	indegree := make([]int, n)
	for _, e := range edges {
		indegree[e[1]]++
	}
	champion := -1
	for i, cnt := range indegree {
		if cnt == 0 {
			if champion == -1 {
				champion = i
			} else {
				// 有多个入度 = 0 的节点
				champion = -1
				break
			}
		}
	}
	return champion
}

func dfsSolution(n int, edges [][]int) int {
	// 1. 构建DAG
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, 0, n)
	}
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
	}
	for i := 0; i < n; i++ {
		// 统计遍历到的节点个数
		visited := make([]bool, n)
		cnt := dfs(adj, i, visited)
		if cnt == n {
			return i
		}
	}
	return -1
}

func dfs(adj [][]int, node int, visited []bool) int {
	var res int
	visited[node] = true
	for _, c := range adj[node] {
		if visited[c] {
			continue
		}
		res += dfs(adj, c, visited)
	}
	return res + 1
}
