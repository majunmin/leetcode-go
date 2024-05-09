package pq_dijkstra

import (
	"container/heap"
	"math"
)

type Graph struct {
	adj []map[int]int
	n   int
}

func Constructor(n int, edges [][]int) Graph {
	adj := make([]map[int]int, n)
	for i := range adj {
		adj[i] = make(map[int]int)
	}
	graph := Graph{
		adj: adj,
		n:   n,
	}
	for _, e := range edges {
		graph.AddEdge(e)
	}
	return graph
}

func (this *Graph) AddEdge(edge []int) {
	this.adj[edge[0]][edge[1]] = edge[2]
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	dists := make([]int, 0)
	for i := 0; i < len(dists); i++ {
		dists[i] = math.MaxInt
	}
	dists[node1] = 0
	//
	pq := NewMinHp()
	pq.offer(pair{node1, 0})
	for pq.Len() > 0 {
		item := pq.poll()
		if item.vertex == node2 {
			return item.cost
		}
		if item.cost > dists[item.vertex] {
			continue
		}
		for next, cost := range this.adj[item.vertex] {
			if item.cost+cost < dists[next] {
				dists[next] = item.cost + cost
				pq.offer(pair{next, item.cost + cost})
			}
		}
	}
	if dists[node2] == math.MaxInt {
		return -1
	}
	return dists[node2]
}

type pair struct {
	vertex int
	cost   int
}

type minHp []pair

func NewMinHp() *minHp {
	pq := new(minHp)
	heap.Init(pq)
	return pq
}

func (m minHp) Len() int {
	return len(m)
}

func (m minHp) Less(i, j int) bool {
	return m[i].cost < m[j].cost
}

func (m minHp) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m *minHp) Push(x any) {
	*m = append(*m, x.(pair))
}

func (m *minHp) Pop() any {
	x := (*m)[len(*m)-1]
	*m = (*m)[:len(*m)-1]
	return x
}

func (m *minHp) offer(p pair) {
	heap.Push(m, p)
}

func (m *minHp) poll() pair {
	return heap.Pop(m).(pair)
}
