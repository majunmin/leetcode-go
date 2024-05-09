package leetcode_0310

import "math"

// https://leetcode.cn/problems/minimum-height-trees/
// 1. 很自然的想法, 枚举 顶点， 然后通过 图的遍历方式(dfs/bfs) 计算树的高度。 类似于题目:
func findMinHeightTrees(n int, edges [][]int) []int {
	if n < 1 {
		return nil
	}
	if n == 1 {
		return []int{0}
	}
	degree := make([]int, n)
	adj := make([][]int, n)
	// 构建出度队列
	for _, e := range edges {
		degree[e[0]]++
		degree[e[1]]++
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	// bfs
	queue := make([]int, 0)
	// 将所有叶子节点入队
	for i := range degree {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}
	var res []int
	for len(queue) > 0 {
		size := len(queue)
		res = res[:0] // 每一层,都要清空 res, 我们要保留的是最后一层的 res
		for i := 0; i < size; i++ {
			item := queue[i]
			res = append(res, item)
			for _, next := range adj[item] {
				degree[next]--
				if degree[next] == 1 {
					queue = append(queue, next)
				}
			}
		}
		queue = queue[size:]
	}
	return res
}

// 枚举根节点,超时😂
func solution1(n int, edges [][]int) []int {
	// param check
	if n < 1 {
		return nil
	}
	//
	var result []int
	//枚举根节点,计算出树的高度
	// 构建图: 邻接数组
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	minHeight := math.MaxInt
	for i := 0; i < n; i++ {
		h := calHeight(adj, -1, i)
		if h > minHeight {
			continue
		}
		if h < minHeight {
			minHeight = h
			result = result[:0]
		}
		result = append(result, i)
	}
	return result
}

// 超时 dfs
// 计算以 parent 为父节点,以 cur 为根节点的树的高度
//func calHeight(adj [][]int, parent int, cur int) int {
//	height := 0
//	for _, child := range adj[cur] {
//		if child == parent {
//			continue
//		}
//		height = max(height, calHeight(adj, cur, child))
//	}
//	return height + 1
//}

func calHeight(adj [][]int, parent int, cur int) int {

	var (
		queue   = make([]int, 0)
		visited = make(map[int]bool)
		cnt     int
	)
	queue = append(queue, cur)
	visited[cur] = true
	for len(queue) > 0 {
		cnt++
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			visited[item] = true
			for _, child := range adj[item] {
				if visited[child] {
					continue
				}
				queue = append(queue, child)
			}
		}
		queue = queue[size:]
	}
	return cnt
}
