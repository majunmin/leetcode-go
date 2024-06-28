package leetcode_100279

import (
	"container/heap"
	"math"
)

type edge struct {
	next int
	w    int
}

func minimumTime2(n int, edges [][]int, disappear []int) []int {
	// param check
	var (
		ans = make([]int, n)
		adj = make([][]edge, n)
	)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{
			next: e[1],
			w:    e[2],
		})
		adj[e[1]] = append(adj[e[1]], edge{
			next: e[0],
			w:    e[2],
		})
	}
	// fill default
	for i := 1; i < n; i++ {
		ans[i] = math.MaxInt
	}
	//计算最短路径
	dfs(adj, disappear, 0, 0, ans)
	for i, cost := range ans {
		if cost == math.MaxInt {
			ans[i] = -1
		}
	}
	return ans
}

func dfs(adj [][]edge, disappear []int, node int, cost int, ans []int) {
	// termiante
	// cost == disappear[node] 也是不可达的
	if cost >= disappear[node] {
		return
	}
	ans[node] = min(ans[node], cost)
	for _, e := range adj[node] {
		dfs(adj, disappear, e.next, cost+e.w, ans)
	}
}

func minimumTime(n int, edges [][]int, disappear []int) []int {
	//flod 算法
	var (
		board = make([][]int, n)
		adj   = make([][]int, n)
		ans   = make([]int, n)
	)
	for i := range board {
		adj[i] = make([]int, n)
		board[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			board[i][j] = 100001
		}
	}
	// fill default
	for i := 1; i < n; i++ {
		ans[i] = math.MaxInt
	}
	for _, e := range edges {
		adj[e[0]][e[1]] = e[2]
		adj[e[1]][e[0]] = e[2]
	}

	//开始 Floyd 算法
	//每个点为中转
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				newDistance := board[i][k] + board[k][j]
				if newDistance < board[i][j] && disappear[j] > board[0][k]+newDistance && disappear[k] > board[0][k] {
					//刷新距离
					board[i][j] = newDistance
				}
			}
		}
	}
	for i := 1; i < n; i++ {
		if board[0][i] == 100001 {
			ans[i] = -1
		} else {
			ans[i] = board[0][i]
		}
	}
	return ans
}

func dijstraSolution(n int, edges [][]int, disappear []int) []int {
	// param check
	var (
		ans    = make([]int, n)
		adj    = make([][]edge, n)
		pQueue = newPQ()
	)
	//构建图
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{
			next: e[1],
			w:    e[2],
		})
		adj[e[1]] = append(adj[e[1]], edge{
			next: e[0],
			w:    e[2],
		})
	}
	for i := 1; i < n; i++ {
		ans[i] = -1
	}
	// 优先级队列优化的 dijstra算法
	pQueue.Push(pair{0, 0})
	for pQueue.Len() > 0 {
		item := pQueue.poll()
		x, dx := item.x, item.distance
		if dx > ans[x] { // x 之前在堆中出现过
			continue
		}
		for _, e := range adj[x] {
			y := e.next
			newDis := dx + e.w
			if newDis < disappear[y] && (ans[y] < 0 || newDis < ans[y]) {
				ans[y] = newDis
				pQueue.offer(pair{y, newDis})
			}
		}
	}
	return ans
}

type pair struct {
	x        int
	distance int
}

type pq []pair

func newPQ() *pq {
	p := new(pq)
	heap.Init(p)
	return p
}

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].distance < p[j].distance
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	*p = append(*p, x.(pair))
}

func (p *pq) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}

func (p *pq) offer(x pair) {
	heap.Push(p, x)
}

func (p *pq) poll() pair {
	return heap.Pop(p).(pair)
}
