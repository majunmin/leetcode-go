package leetcode_0621

import (
	"container/heap"
	"slices"
)

func leastInterval(tasks []byte, n int) int {
	// 一种巧妙的解法
	var (
		cnts = make([]int, 26)
		tail int
	)
	for _, item := range tasks {
		cnts[item-'A']++
	}
	slices.SortFunc(cnts, func(a, b int) int {
		return b - a
	})
	maxFreq := cnts[0]
	for i := 0; i < len(cnts); i++ {
		if cnts[i] == 0 {
			break
		}
		if cnts[i] == maxFreq {
			tail++
		}
	}
	return min((maxFreq-1)*(n+1)+tail, len(tasks))
}

// 模拟的方式
func solution1(tasks []byte, n int) int {
	var (
		typeCnts  = make(map[int]int)
		priorityQ = newPq()
		result    int
	)
	for _, item := range tasks {
		typeCnts[int(item-'A')]++
	}
	for _, cnt := range typeCnts {
		heap.Push(priorityQ, cnt)
	}
	n += 1
	for priorityQ.Len() > 0 {
		temp := make([]int, 0)
		k := min(n, priorityQ.Len())
		for i := 0; i < k; i++ {
			item := heap.Pop(priorityQ).(int) - 1
			if item != 0 {
				temp = append(temp, item)
			}
		}
		if len(temp) == 0 {
			//最后一轮
			result += k
		} else {
			result += n
		}
		for _, item := range temp {
			heap.Push(priorityQ, item)
		}
	}
	return result
}

type pq []int

func newPq() *pq {
	p := new(pq)
	return p
}

func (p pq) Len() int {
	return len(p)
}

func (p pq) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p pq) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *pq) Push(x any) {
	*p = append(*p, x.(int))
}

func (p *pq) Pop() any {
	x := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return x
}
