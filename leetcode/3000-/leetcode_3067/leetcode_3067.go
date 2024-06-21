package leetcode_3067

// https://leetcode.cn/problems/count-pairs-of-connectable-servers-in-a-weighted-tree-network/
func countPairsOfConnectableServers(edges [][]int, signalSpeed int) []int {
	// 构建 树
	var (
		result = make([]int, 0)
		graph  = make(map[int][][2]int, 0)
		n      int // n = len(edges) + 1
	)

	for _, e := range edges {
		start, end, w := e[0], e[1], e[2]
		n = max(n, e[0], e[1])
		graph[start] = append(graph[start], [2]int{end, w})
		graph[end] = append(graph[end], [2]int{start, w})
	}
	// 2. 枚举每一个节点作为根节点, 将根节点
	// 统计每颗子树 符合条件的节点数
	for i := 0; i <= n; i++ {
		var res, sum int
		for _, e := range graph[i] {
			cnt := dfs(i, e[0], e[1], signalSpeed, graph)
			res += cnt * sum
			sum += cnt
		}
		result = append(result, res)
	}
	return result
}

func dfs(prev, cur, w, signalSpeed int, graph map[int][][2]int) int {
	var res int
	if w%signalSpeed == 0 {
		res++
	}
	for _, e := range graph[cur] {
		if e[0] == prev {
			continue
		}
		res += dfs(cur, e[0], w+e[1], signalSpeed, graph)
	}
	return res
}
