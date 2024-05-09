package leetcode_1094

import (
	"container/heap"
	"math"
	"sort"
)

// https://leetcode.cn/problems/car-pooling/
func carPooling(trips [][]int, capacity int) bool {
	// param check
	minVal, maxVal := math.MaxInt, 0
	cntMap := make(map[int]int)
	for _, path := range trips {
		weight, from, to := path[0], path[1], path[2]
		minVal = min(minVal, from)
		maxVal = max(maxVal, to)
		for i := from; i < to; i++ {
			cntMap[i] += weight
		}
	}
	for i := minVal; i <= maxVal; i++ {
		if cntMap[i] > capacity {
			return false
		}
	}
	return true
}

// 差分数组
func solution1(trips [][]int, capacity int) bool {
	sites := make([]int, 1001)
	for _, path := range trips {
		weight, from, to := path[0], path[1], path[2]
		sites[from] += weight // 上车
		sites[to] -= weight   // 下车
	}
	var total int
	for _, site := range sites {
		total += site
		if total > capacity {
			return false
		}
	}
	return true
}

// solution2 小顶堆
type hp struct {
	trips [][]int
}

func NewHeap() *hp {
	return &hp{}
}

func (h *hp) Len() int {
	return len(h.trips)
}

func (h *hp) Less(i, j int) bool {
	return h.trips[i][2] < h.trips[j][2]
}

func (h *hp) Swap(i, j int) {
	h.trips[i], h.trips[j] = h.trips[j], h.trips[i]
}

func (h *hp) Push(x any) {
	h.trips = append(h.trips, x.([]int))
}

func (h *hp) Pop() any {
	item := h.trips[len(h.trips)-1]
	h.trips = h.trips[:len(h.trips)-1]
	return item
}

func solution2(trips [][]int, capacity int) bool {
	minHeap := NewHeap()

	sort.Slice(trips, func(i, j int) bool {
		return trips[i][1] < trips[j][1]
	})

	for i := 0; i < len(trips); i++ {
		weight, from, _ := trips[i][0], trips[i][1], trips[i][2]
		for minHeap.Len() > 0 && minHeap.trips[0][2] <= from {
			path := heap.Pop(minHeap).([]int)
			capacity += path[0]
		}

		capacity -= weight
		if capacity < 0 {
			return false
		}
		heap.Push(minHeap, trips[i])
	}
	return true
}
