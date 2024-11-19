package leetcode_3249

func countGoodNodes(edges [][]int) int {
	// param check
	if len(edges) == 0 {
		return 0
	}
	var (
		n      = len(edges) + 1
		adj    = make([][]int, n)
		result int

		// 统计当前节点及其子节点的数目
		dfs func(cur, parent int) int
	)

	dfs = func(cur, parent int) int {
		var cnt int
		var prevCnt int
		isGoodNode := true
		for _, child := range adj[cur] {
			if child == parent {
				continue
			}
			curCnt := dfs(child, cur)
			if prevCnt != 0 && prevCnt != curCnt {
				isGoodNode = false
			}
			cnt += curCnt
			prevCnt = curCnt
		}
		if isGoodNode {
			result++
		}
		return cnt + 1
	}

	// 树的根节点是 0
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	dfs(0, -1)
	return result
}
