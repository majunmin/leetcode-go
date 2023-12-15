package lcr_0168

import (
	"container/heap"
	"sort"
)

type hp struct {
	sort.IntSlice
}

func newHeap() *hp {
	return &hp{}
}

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	x := h.IntSlice[len(h.IntSlice)] - 1
	h.IntSlice = h.IntSlice[:len(h.IntSlice)-1]
	return x
}

func nthUglyNumber(n int) int {
	// param check
	if n < 1 {
		panic("invalid param")
	}
	minHeap := newHeap()
	visited := make(map[int]bool)
	visited[1] = true
	heap.Push(minHeap, 1)
	for i := 1; i < n; i++ {
		item := heap.Pop(minHeap).(int)
		if !visited[2*item] {
			heap.Push(minHeap, 2*item)
			visited[2*item] = true
		}
		if !visited[3*item] {
			heap.Push(minHeap, 3*item)
			visited[3*item] = true
		}
		if !visited[5*item] {
			heap.Push(minHeap, 5*item)
			visited[5*item] = true
		}
	}
	return heap.Pop(minHeap).(int)
}
