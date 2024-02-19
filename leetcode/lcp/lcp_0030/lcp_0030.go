package lcp_0030

import (
	"container/heap"
	"sort"
)

type minHeap struct {
	sort.IntSlice
}

func newMinHeap() *minHeap {
	minHp := &minHeap{IntSlice: make([]int, 0)}
	heap.Init(minHp)
	return minHp
}

func (m *minHeap) Push(x any) {
	m.IntSlice = append(m.IntSlice, x.(int))
}

func (m *minHeap) Pop() any {
	x := m.IntSlice[len(m.IntSlice)-1]
	m.IntSlice = m.IntSlice[:len(m.IntSlice)-1]
	return x
}
func (m *minHeap) push(x any) {
	heap.Push(m, x)
}

func (m *minHeap) pop() any {
	return heap.Pop(m)
}

// https://leetcode.cn/problems/p0NxJO/
func magicTower(nums []int) int {
	minHp := newMinHeap()
	var result int // 交换次数
	tailSum := 0
	leftSum, rightSum := 1, 0
	for _, num := range nums {
		if num < 0 {
			minHp.push(num)
			rightSum += num
		} else {
			leftSum += num
		}
		for leftSum+rightSum <= 0 && minHp.IntSlice.Len() > 0 {
			item := minHp.pop().(int)
			tailSum += item
			rightSum -= item
			result++
		}
		// process
	}
	if leftSum+rightSum+tailSum <= 0 {
		return -1
	}
	return result
}
