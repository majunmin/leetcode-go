package leetcode_2191

import "sort"

// https://leetcode.cn/problems/all-ancestors-of-a-node-in-a-directed-acyclic-graph/?envType=daily-question&envId=2024-04-05
func getAncestors(n int, edges [][]int) [][]int {
	// dfs
	var (
		parent   = make([]map[int]bool, n)
		inDegree = make([]int, n)
		adj      = make([][]int, n)
		queue    = make([]int, 0)
	)
	// 构建DAG图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		inDegree[e[1]]++

	}
	// dfs 遍历图， 根节点?
	//bfs 找到入度为0 的节点为根节点
	//
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
		parent[i] = make(map[int]bool)
	}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			for _, next := range adj[cur] {
				parent[next][cur] = true
				for p, _ := range parent[cur] {
					parent[next][p] = true
				}
				inDegree[next]--
				if inDegree[next] == 0 {
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	return conver2Res(parent)
}

func conver2Res(parent []map[int]bool) [][]int {
	result := make([][]int, len(parent))
	for i := range parent {
		arr := make([]int, 0, len(parent[i]))
		for p := range parent[i] {
			arr = append(arr, p)
		}
		sort.Ints(arr)
		result[i] = arr
	}
	return result
}
