package leetcode_1976

import (
	"container/heap"
	"math"
)

func countPaths(n int, roads [][]int) int {
	var (
		adj = make([][]edge, n)
		pq  = newPriorityQueue()
		//visited = make([]bool, n) // 表示节点i 是否添加到优先级队列中过
		dis = make([]int, n) // 表示节点i 到起始节点的最短距离
		dp  = make([]int, n) // 表示节点i 到起始节点的路径数
	)
	// init
	for i := 0; i < n; i++ {
		adj[i] = make([]edge, n)
		dis[i] = math.MaxInt
	}
	for _, r := range roads {
		s, t, w := r[0], r[1], r[2]
		adj[s] = append(adj[s], edge{s: s, t: t, cost: w})
		adj[t] = append(adj[t], edge{s: t, t: s, cost: w})
	}
	//起始节点到自己的 最短距离是 0
	dis[0] = 0
	dp[0] = 1
	heap.Init(pq)
	pq.push(vertex{t: 0, cost: 0})
	//visited[0] = true
	for pq.Len() > 0 {
		prev := pq.pop().(vertex)
		u := prev.t
		t := prev.cost
		// 剪枝
		if t > dis[u] {
			continue
		}

		for _, e := range adj[u] {
			v, w := e.t, e.cost
			if t+w < dis[v] {
				dis[v] = t + w
				dp[v] = dp[u]
				pq.push(vertex{t: v, cost: t + w})
				//visited[v] = true
			} else if w+w == dis[v] {
				dp[v] = (dp[u] + dp[v]) % (1e9 + 7)
			}
		}
	}
	return dp[n-1]
}

type edge struct {
	s    int // 起始节点id
	t    int // 目标节点id
	cost int // 起点s 到 终点t 路径长度
}

// 记录目标节点到起点的最短路径
type vertex struct {
	t    int
	cost int
}

type priorityQueue []vertex

func newPriorityQueue() *priorityQueue {
	pq := &priorityQueue{}
	return pq
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
