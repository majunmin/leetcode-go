package lcr_0159

import (
	"container/heap"
	"sort"
)

type minHeap struct {
	sort.IntSlice
}

func newMinHeap() *minHeap {
	return &minHeap{}
}

func (m *minHeap) Push(x interface{}) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *minHeap) Pop() interface{} {
	length := len(m.IntSlice)
	x := m.IntSlice[length-1]
	m.IntSlice = m.IntSlice[:length-1]
	return x
}

func (m *minHeap) Less(i, j int) bool {
	return m.IntSlice[i] < m.IntSlice[j]
}

func inventoryManagement(stock []int, cnt int) []int {
	mh := newMinHeap()
	heap.Init(mh)
	for _, item := range stock {
		heap.Push(mh, item)
	}

	//
	result := make([]int, 0, cnt)
	for i := 0; i < cnt; i++ {
		item := heap.Pop(mh)
		result = append(result, item.(int))
	}
	return result
}
