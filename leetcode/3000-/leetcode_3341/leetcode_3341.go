package leetcode_3341

import (
	"container/heap"
	"math"
)

var (
	dirtions = [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/find-minimum-time-to-reach-last-room-i/description/?envType=daily-question&envId=2025-05-07
func minTimeToReach(moveTime [][]int) int {
	var (
		m, n = len(moveTime), len(moveTime[0])

		d         = make([][]int, m)
		visited   = make([][]bool, m)
		priorityQ = newPq()
	)

	// init state
	for i := 0; i < m; i++ {
		d[i] = make([]int, n)
		visited[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			d[i][j] = math.MaxInt32
		}
	}
	d[0][0] = 0
	//
	heap.Push(priorityQ, &state{0, 0, 0})
	for priorityQ.Len() > 0 {
		item := heap.Pop(priorityQ).(*state)
		if visited[item.x][item.y] {
			continue
		}
		visited[item.x][item.y] = true

		for _, dir := range dirtions {
			nx, ny := item.x+dir[0], item.y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			dist := max(d[item.x][item.y], moveTime[nx][ny]) + 1
			if d[nx][ny] > dist {
				d[nx][ny] = dist
				heap.Push(priorityQ, &state{nx, ny, dist})
			}
		}
	}
	return d[m-1][n-1]
}

type state struct {
	x, y int
	cost int
}

type pq []*state

func newPq() *pq {
	q := make(pq, 0)
	heap.Init(&q)
	return &q
}

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i].cost < p[j].cost
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	*p = append(*p, x.(*state))
}

func (p *pq) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}
