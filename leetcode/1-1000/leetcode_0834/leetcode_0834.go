package leetcode_0834

// https://leetcode.cn/problems/sum-of-distances-in-tree/
// 题解:
// https://leetcode.cn/problems/sum-of-distances-in-tree/solutions/2345592/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/
func sumOfDistancesInTree(n int, edges [][]int) []int {
	// 换根dp
	var (
		size   = make([]int, n)
		result = make([]int, n)
		adj    = make([][]int, n)
	)
	// 构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	dfs(adj, size, result, 0, -1, 0)
	reroot(adj, size, result, 0, -1)
	return result
}

func reroot(adj [][]int, size []int, result []int, cur int, parent int) {
	for _, next := range adj[cur] {
		if next == parent {
			continue
		}
		result[next] = result[cur] + len(adj) - 2*size[next]
		reroot(adj, size, result, next, cur)
	}
}

func dfs(adj [][]int, size []int, result []int, cur int, parent int, depth int) {
	result[0] += depth
	var curSize int
	for _, next := range adj[cur] {
		if next == parent {
			continue
		}
		dfs(adj, size, result, next, cur, depth+1)
		curSize += size[next]
	}
	size[cur] = curSize + 1
}
