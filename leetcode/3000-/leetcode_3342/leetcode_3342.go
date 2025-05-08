package leetcode_3342

import (
	"container/heap"
	"math"
	"slices"
)

var (
	directions = [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

// https://leetcode.cn/problems/find-minimum-time-to-reach-last-room-ii/?envType=daily-question&envId=2025-05-08
func minTimeToReach(moveTime [][]int) int {
	var (
		m, n = len(moveTime), len(moveTime[0])
		pq   = newPriorityQueue()
	)
	visited := make([][]bool, m)
	dp := make([][]int, m)
	// init state
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
		dp[i] = slices.Repeat([]int{math.MaxInt - 1}, n)
	}
	dp[0][0] = 0
	heap.Push(pq, &state{0, 0, 0})
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*state)
		x, y := item.x, item.y

		if visited[x][y] {
			continue
		}
		visited[x][y] = true

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx < 0 || nx >= m || ny < 0 || ny >= n {
				continue
			}
			dist := max(moveTime[nx][ny], dp[x][y]) + 1 + (x+y)%2
			if dp[nx][ny] > dist {
				dp[nx][ny] = dist
				heap.Push(pq, &state{nx, ny, dist})
			}
		}
	}

	return dp[m-1][n-1]
}

// dijstra 算法

type state struct {
	x, y int
	cost int
}

type priorityQueue []*state

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
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(x any) {
	*p = append(*p, x.(*state))
}

func (p *priorityQueue) Pop() any {
	x := (*p)[p.Len()-1]
	*p = (*p)[:p.Len()-1]
	return x
}
