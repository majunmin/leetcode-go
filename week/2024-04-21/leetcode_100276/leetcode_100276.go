package leetcode_100276

import (
	"container/heap"
	"math"
)

// Dijkstra 算法
func findAnswer(n int, edges [][]int) []bool {
	var (
		adj      = make([][]edge, n)
		pq       = NewPriorityQueue()
		vertexs  = make([]*vertex, n)
		result   = make([]bool, len(edges))
		edge2Idx = make(map[int]int)
	)
	// init adj
	for i, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{
			from:   e[0],
			to:     e[1],
			weight: e[2],
		})
		adj[e[1]] = append(adj[e[1]], edge{
			from:   e[1],
			to:     e[0],
			weight: e[2],
		})
		edge2Idx[e[0]<<20|e[1]] = i
		edge2Idx[e[1]<<20|e[0]] = i
	}
	for i := 0; i < n; i++ {
		vertexs[i] = &vertex{
			idx:       i,
			minCost:   math.MaxInt,
			cost2Path: make(map[int]bool),
		}
	}
	vertexs[0].minCost = 0
	pq.offer(vertexs[0])
	for pq.Len() > 0 {
		from := pq.poll()
		for _, next := range adj[from.idx] {
			to := vertexs[next.to]
			if from.minCost+next.weight <= to.minCost {
				// add path = to.cost2Path
				to.minCost = from.minCost + next.weight
				for k := range from.cost2Path {
					to.cost2Path[k] = true
				}
				to.cost2Path[from.idx<<20|to.idx] = true
				pq.offer(to)
			}
		}
	}

	target := vertexs[n-1]
	for k := range target.cost2Path {
		result[edge2Idx[k]] = true
	}
	return result
}

type vertex struct {
	idx       int
	minCost   int // 从 0 到该节点的最小开销
	cost2Path map[int]bool
}

type edge struct {
	from, to int
	weight   int
}

type priorityQueue []*vertex

func NewPriorityQueue() *priorityQueue {
	pq := new(priorityQueue)
	heap.Init(pq)
	return pq
}

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].minCost < p[j].minCost
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(*vertex))
}

func (p *priorityQueue) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func (p *priorityQueue) offer(x *vertex) {
	heap.Push(p, x)
}

func (p *priorityQueue) poll() *vertex {
	return heap.Pop(p).(*vertex)
}
