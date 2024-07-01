package leetcode_2065

import (
	"container/heap"
	"math"
)

type edge struct {
	to   int
	time int // cost
}

// https://leetcode.cn/problems/maximum-path-quality-of-a-graph/?envType=daily-question&envId=2024-07-01
// solution1 暴力搜索
func maximalPathQuality(values []int, edges [][]int, maxTime int) int {
	// 构建图
	// param check
	if len(values) == 0 {
		return 0
	}

	var (
		result  int
		adj     = make([][]edge, len(values))
		visited = make([]bool, len(values))
		dfs     func(int, int, int)
	)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], edge{e[0], e[2]})
	}
	dfs = func(cur int, leftTime int, sumValue int) {
		if leftTime < 0 {
			return
		}
		if cur == 0 {
			// update result
			result = max(result, sumValue)
		}
		for _, edge := range adj[cur] {
			next := edge.to
			cost := edge.time

			if visited[next] {
				dfs(next, leftTime-cost, sumValue)
			} else {
				visited[next] = true
				dfs(next, leftTime-cost, sumValue+values[next])
				visited[next] = false
			}
		}
	}

	visited[0] = true
	// dfs
	dfs(0, maxTime, values[0])
	return result
}

// solution2 dijkstra算法
func maximalPathQuality2(values []int, edges [][]int, maxTime int) int {
	if len(values) == 0 {
		return 0
	}
	var (
		n        = len(values)
		distance = make([]int, len(values))
		adj      = make([][]edge, n)
		minHp    = new(hp)
	)
	for i := 1; i < n; i++ {
		distance[i] = math.MaxInt
	}
	// 构建邻接图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
		adj[e[1]] = append(adj[e[1]], edge{e[0], e[2]})
	}
	heap.Init(minHp)
	heap.Push(minHp, pair{0, 0})
	// Dijkstra
	// 计算最短路径 用于剪枝
	for minHp.Len() > 0 {
		p := heap.Pop(minHp).(pair)
		x := p.x
		dx := p.dis
		// 长路径
		if dx > distance[x] {
			continue
		}
		for _, e := range adj[x] {
			y := e.to
			cost := e.time
			newDis := dx + cost
			if newDis < distance[y] {
				distance[y] = newDis
				heap.Push(minHp, pair{newDis, y})
			}
		}
	}
	var (
		visited = make([]bool, n)
		dfs     func(int, int, int)
		result  int
	)
	dfs = func(cur int, leftTime int, sumValue int) {
		if leftTime < 0 {
			return
		}
		if cur == 0 {
			// update result
			result = max(result, sumValue)
		}
		for _, e := range adj[cur] {
			next, cost := e.to, e.time
			// 距离不够返回 0 的了， 剪枝
			if leftTime < cost+distance[next] {
				continue
			}
			if visited[next] {
				dfs(next, leftTime-cost, sumValue)
			} else {
				visited[next] = true
				dfs(next, leftTime-cost, sumValue+values[next])
				visited[next] = false
			}
		}
	}

	visited[0] = true
	dfs(0, maxTime, values[0])
	return result
}

type pair struct {
	dis int
	x   int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}
func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}
