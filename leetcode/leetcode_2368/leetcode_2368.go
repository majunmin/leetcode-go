package leetcode_2368

import "fmt"

// https://leetcode.cn/problems/reachable-nodes-with-restrictions/
func reachableNodes(n int, edges [][]int, restricted []int) int {
	return ufSolution(n, edges, restricted)
}

// 并查集解法
func ufSolution(n int, edges [][]int, restricted []int) int {
	var (
		uf          = NewUnionFind(n)
		restrictSet = make([]bool, n)
	)
	for _, item := range restricted {
		restrictSet[item] = true
	}
	for _, e := range edges {
		if restrictSet[e[0]] || restrictSet[e[1]] {
			continue
		}
		uf.union(e[0], e[1])
	}
	return uf.count()
}

func bfsSolution(n int, edges [][]int, restricted []int) int {
	// 构建图 + visitedSet
	var (
		adj     = make([][]int, n)
		visited = make([]bool, n)
	)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	for _, item := range restricted {
		visited[item] = true
	}
	//
	var cnt int
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			item := queue[i]
			if visited[item] {
				continue
			}
			visited[item] = true
			cnt++
			for _, next := range adj[item] {
				queue = append(queue, next)
			}
		}
		queue = queue[size:]
	}
	return cnt
}

func dfsSolution(n int, edges [][]int, restricted []int) int {
	// param check
	if n < 1 {
		return 0
	}
	// 构建无向图
	var (
		adj         = make([][]int, n)
		restrictSet = make(map[int]bool)
	)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	for _, item := range restricted {
		restrictSet[item] = true
	}

	var cnt int
	var dfs func(cur, parent int)
	dfs = func(cur, parent int) {
		// terminate
		if restrictSet[cur] {
			return
		}

		cnt++
		for _, next := range adj[cur] {
			// 避免重复遍历
			if next == parent {
				continue
			}
			dfs(next, cur)
		}
	}
	return cnt
}

type UnionFind struct {
	p []int
	//cnt int
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &UnionFind{p: p}
}

func (uf *UnionFind) find(x int) int {
	//if x != uf.p[x] {
	//	uf.find(uf.p[x])
	//}
	//return x
	for x != uf.p[x] {
		// 压缩
		x = uf.p[x]
		uf.p[x] = uf.p[uf.p[x]]
	}
	return uf.p[x]
}

func (uf *UnionFind) union(x, y int) {
	rootx, rooty := uf.find(x), uf.find(y)
	if rootx == rooty {
		return
	}
	uf.p[rootx] = rooty
}

func (uf *UnionFind) count() int {
	var cnt int
	root := uf.find(0)
	for i := 0; i < len(uf.p); i++ {
		if uf.find(i) == root {
			cnt++
			continue
		}
		fmt.Print(i)
	}
	return cnt
}
