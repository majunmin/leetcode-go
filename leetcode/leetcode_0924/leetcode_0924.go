package leetcode_0924

import "sort"

// https://leetcode.cn/problems/minimize-malware-spread/?envType=daily-question&envId=2024-04-16
func minMalwareSpread(graph [][]int, initial []int) int {
	// param check
	if len(graph) == 0 || len(graph[0]) == 0 {
		return -1
	}
	if len(initial) == 0 {
		return -1
	}
	sort.Ints(initial)
	var (
		n   = len(graph)
		uf  = NewUnionFind(n)
		adj = make([][]int, n)
	)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				uf.union(i, j)
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	//
	var (
		// 统计每个连通分量的 感染节点数
		inject = make([]int, n)
	)

	for _, e := range initial {
		inject[uf.find(e)]++
	}
	var (
		maxCnt int // 最大感染的节点数
		node   = initial[0]
	)

	for _, e := range initial {
		if inject[uf.find(e)] > 1 {
			continue
		}
		if uf.size[uf.find(e)] > maxCnt {
			maxCnt = uf.size[uf.find(e)]
			node = e
		}
	}
	return node
}

type UnionFind struct {
	p     []int
	count int
	rank  []int
	size  []int
}

func NewUnionFind(cap int) *UnionFind {
	parent := make([]int, cap)
	sizes := make([]int, cap)
	for i := 0; i < cap; i++ {
		parent[i] = i
		sizes[i] = 1
	}
	return &UnionFind{
		p:    parent,
		rank: make([]int, cap),
		size: sizes,
	}
}

func (uf *UnionFind) union(i, j int) {
	p1, p2 := uf.find(i), uf.find(j)
	if p1 == p2 {
		return
	}
	if uf.rank[p1] < uf.rank[p2] {
		p1, p2 = p2, p1
	}
	uf.size[p1] += uf.size[p2]
	uf.size[p2] = 0
	uf.p[p2] = p1
	if uf.rank[p1] == uf.rank[p2] {
		uf.rank[p1]++
	}
}

func (uf *UnionFind) Count() int {
	return uf.count
}

func (uf *UnionFind) find(i int) int {
	if i != uf.p[i] {
		uf.p[i] = uf.find(uf.p[i])
	}
	return uf.p[i]
}

func minMalwareSpread2(graph [][]int, initial []int) int {
	var (
		n        = len(graph)
		adj      = make([][]int, n)
		initials = make(map[int]bool)
		visited  = make(map[int]bool)
	)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				adj[i] = append(adj[i], j)
				adj[j] = append(adj[j], i)
			}
		}
	}
	sort.Ints(initial)
	var (
		nodeId int
		cnt    int
		dfs    func(cur int)
		result = initial[0]
		maxCnt int
	)

	dfs = func(cur int) {
		if visited[cur] {
			return
		}
		cnt++
		visited[cur] = true
		if initials[cur] {
			if nodeId == -1 {
				nodeId = cur
			} else if nodeId > 0 {
				nodeId = -2
			}
		}
		for _, next := range adj[cur] {
			dfs(next)
		}
	}

	for i := 0; i < n; i++ {
		if visited[i] {
			continue
		}
		nodeId = -1
		cnt = 0
		dfs(i)
		if nodeId > 0 && cnt > maxCnt {
			result = nodeId
			cnt = maxCnt
		}
	}
	return result
}
