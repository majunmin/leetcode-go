package dijkstra

import (
	"container/heap"
	"math"
)

// dijkstra 实现
// 顶点定义
type vertex struct {
	tid  int //顶点id
	idx  int // 在优先级队列的索引的位置
	cost int // 顶点t到起点s的最短路径
}

type edge struct {
	sid int // 起始顶点id
	tid int // 终点id
	w   int // 权重
}

type Graph struct {
	// 图的邻接表表示
	adj  [][]edge
	size int
}

func NewGraph(size int) *Graph {
	adj := make([][]edge, size)
	for i := 0; i < size; i++ {
		adj[i] = make([]edge, 0)
	}
	return &Graph{
		size: size,
		adj:  adj,
	}
}

func (g *Graph) addEdge(s, t, w int) {
	g.adj[s] = append(g.adj[s], edge{s, t, w})
}

// 计算 s t之间的最短路径
func (g *Graph) dijkstra(s, t int) int {
	var (
		predecessor = make([]int, g.size) // 用于输出单源最短路径
		vertexs     = make([]vertex, g.size)
		visited     = make([]bool, g.size)
		pq          = newPriorityQueue()
	)

	for i := 0; i < g.size; i++ {
		vertexs[i] = vertex{tid: i, idx: i, cost: math.MaxInt}
	}
	vertexs[s].cost = 0
	pq.push(vertex{tid: s, idx: s, cost: 0})
	visited[0] = true
	for pq.Len() > 0 {
		minVertex := pq.pop().(vertex)
		// terminate
		if minVertex.tid == t {
			break
		}
		//
		for _, e := range g.adj[minVertex.tid] {
			nexttId, w := e.tid, e.w
			nextVertex := vertexs[nexttId]
			if minVertex.cost+w < nextVertex.cost {
				nextVertex.cost = minVertex.cost + w
				predecessor[nexttId] = minVertex.tid
				if visited[nexttId] {
					heap.Fix(pq, nextVertex.idx) // 更新优先级队列中的值
				} else {
					pq.push(nextVertex)
					visited[nexttId] = true
				}
			}
		}
	}
	return vertexs[t].cost
}

// 邻接表实现
// 小顶堆
type priorityQueue []vertex

func newPriorityQueue() *priorityQueue {
	pq := &priorityQueue{}
	heap.Init(pq)
	return pq
}

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p priorityQueue) Swap(i, j int) {
	p[i].idx, p[j].idx = p[j].idx, p[i].idx
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(vertex))
}

func (p *priorityQueue) Pop() any {
	item := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return item
}

func (p *priorityQueue) push(x any) {
	heap.Push(p, x)
}

func (p *priorityQueue) pop() any {
	return heap.Pop(p)
}
