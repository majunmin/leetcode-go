package leetcode_3112

import "container/heap"

// https://leetcode.cn/problems/minimum-time-to-visit-disappearing-nodes/?envType=daily-question&envId=2024-07-18
// Dijkstra 算法
func minimumTime(n int, edges [][]int, disappear []int) []int {
	var (
		adj    = make([][]item, n)
		pq     = newPQ()
		result = make([]int, n)
	)

	// init
	for i := 1; i < n; i++ {
		result[i] = -1
	}
	// 构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], item{to: e[1], cost: e[2]})
		adj[e[1]] = append(adj[e[1]], item{to: e[0], cost: e[2]})
	}
	heap.Push(pq, item{to: 0, cost: 0})
	for pq.Len() > 0 {
		e := heap.Pop(pq).(item)
		u, w := e.to, e.cost
		if w > result[u] { // u 之前出堆过
			continue
		}
		for _, edge := range adj[u] {
			v, cost := edge.to, edge.cost
			if w+cost < disappear[v] && (result[v] == -1 || w+cost < result[v]) {
				heap.Push(pq, item{to: v, cost: w + cost})
				result[v] = w + cost
			}
		}
	}

	return result
}

type item struct {
	to   int
	cost int
}
type priorityQueue []item

func newPQ() *priorityQueue {
	p := new(priorityQueue)
	heap.Init(p)
	return p
}

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(item))
}

func (p *priorityQueue) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}
