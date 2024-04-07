package leetcode_100244

func minimumCost(n int, edges [][]int, query [][]int) []int {
	// dfsSolution()
	//buildGraph
	graph := make(map[int][][2]int)
	for _, e := range edges {
		graph[e[0]] = append(graph[e[0]], [2]int{e[1], e[2]})
		graph[e[1]] = append(graph[e[1]], [2]int{e[0], e[2]})
	}
	// fill -1, 分组用
	ids := make([]int, n)
	for i := 0; i < len(ids); i++ {
		ids[i] = -1
	}
	// 记录每个联通快的 边权and
	ccAnd := make([]int, 0)
	for i := 0; i < len(ids); i++ {
		// 没有访问过
		if ids[i] < 0 {
			ccAnd = append(ccAnd, dfs(graph, len(ccAnd), i, ids))
		}
	}
	// query
	var result []int
	for _, item := range query {
		var res int
		if item[0] == item[1] {
			res = 0
		} else if ids[item[0]] != ids[item[1]] {
			res = -1
		} else {
			res = ccAnd[ids[item[0]]]
		}
		result = append(result, res)
	}
	return result
}

func dfs(g map[int][][2]int, curId, i int, ids []int) int {
	ids[i] = curId
	and := -1 // 全1
	for _, e := range g[i] {
		and &= e[1]
		if ids[e[0]] < 0 { // 没被访问过
			and &= dfs(g, curId, e[0], ids)
		}
	}
	return and
}
