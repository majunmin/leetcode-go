package leetcode_1976

import (
	"container/heap"
	"math"
)

// https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/?envType=daily-question&envId=2024-03-05
func countPaths(n int, roads [][]int) int {
	// 带权重的无向图 构建
	// 朴素dijkstra 算法,
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < len(adj[i]); j++ {
			adj[i][j] = math.MinInt >> 1 // 防止溢出
		}
	}
	for _, items := range roads {
		adj[items[0]][items[1]] = items[2]
		adj[items[1]][items[0]] = items[2]
	}
	var (
		// 表示 节点0 到节点 i的最短距离
		dis = make([]int, n)
		// 表示节点0 到节点 i 的路径数目
		dp   = make([]int, n)
		done = make([]bool, n)
	)
	// init state
	dp[0] = 1
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt
	}
	//
	for {
		s := -1
		for i, ok := range done {
			if !ok && (s < 0 || dis[i] > dis[s]) {
				s = i
			}
		}
		// 起点 == n-1
		//不可能找到比dis[n-1] 更短,或者一样短的路径了(边的权重都是整数)
		if s == n-1 {
			return dp[n-1]
		}
		done[s] = true             // / 最短路长度已确定（无法变得更小）
		for t, w := range adj[s] { // 尝试更新 x 的邻居的最短路
			if dis[s]+w < dis[t] {
				dis[t] = dis[s] + w
				dp[t] = dp[s]
			} else if dis[s]+w == dis[t] {
				// 和之前路径一样长
				dp[t] = (dp[s] + dp[t]) % (1_000_000_007)
			}
		}
	}
}

// dijstela 算法
// floyd 算法

// 邻接表 + 优先级队列实现的方式

func countPaths2(n int, roads [][]int) int {
	// 构建邻接表, 适合与稀疏表
	var (
		adj = make([][]edge, n)
		dis = make([]int, n)
		dp  = make([]int, n)
	)

	for i := 0; i < n; i++ {
		adj[i] = make([]edge, n)
	}
	for _, road := range roads {
		s, t, w := road[0], road[1], road[2]
		adj[s] = append(adj[s], edge{t, w})
		adj[t] = append(adj[t], edge{s, w})
	}
	var ()
	// init state
	for i := 1; i < n; i++ {
		dis[i] = math.MaxInt
	}
	dp[1] = 1
	// process
	pq := priorityQueue{}
	heap.Init(&pq)
	pq.push(pair{0, 0})
	for len(pq) > 0 {
		p := pq.pop().(pair)
		t, u := p.first, p.second
		if t > dis[u] {
			continue
		}
		for _, e := range adj[u] {
			v, w := e.to, e.cost
			if t+w < dis[v] {
				dis[v] = t + w
				dp[v] = dp[u]
				pq.push(pair{t + w, v})
			} else if t+w == dis[v] {
				dp[v] = (dp[u] + dp[v]) % (1_000_000_007)
			}
		}
	}
	return dp[n-1]
}

type edge struct {
	to   int
	cost int
}

type pair struct {
	first, second int
}

type priorityQueue []pair

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].first < p[j].first
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(pair))
}

func (p *priorityQueue) Pop() any {
	item := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return item
}

func (p *priorityQueue) pop() any {
	return heap.Pop(p)
}
func (p *priorityQueue) push(x any) {
	heap.Push(p, x)
}
